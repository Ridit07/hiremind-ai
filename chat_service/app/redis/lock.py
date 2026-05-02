from app.redis.client import redis_client


def acquire_lock(key, timeout=5, blocking_timeout=2):
    lock_key = f"lock:{key}"

    lock = redis_client.lock(
        name=lock_key,
        timeout=timeout,               
        blocking=True,                 
        blocking_timeout=blocking_timeout
    )

    acquired = lock.acquire()

    if not acquired:
        return None

    return lock  


def release_lock(lock):
    if lock:
        try:
            lock.release()
        except Exception:
            pass  