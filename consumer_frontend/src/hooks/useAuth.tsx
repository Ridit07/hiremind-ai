'use client';

import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  useCallback,
  ReactNode,
} from 'react';
import { apiClient, LoginResponse, SignupResponse } from '@/services/api';

interface AuthContextType {
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
  accessToken: string | null;
  refreshToken: string | null;
  signup: (email: string, password: string, userType: string, phoneNumber?: string) => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  clearError: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

const TOKEN_STORAGE_KEY = 'auth_tokens';
const REFRESH_TOKEN_KEY = 'refresh_token';

interface StoredTokens {
  accessToken: string;
  refreshToken: string;
  expiresAt: number;
}

export function AuthProvider({ children }: { children: ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [accessToken, setAccessToken] = useState<string | null>(null);
  const [refreshToken, setRefreshToken] = useState<string | null>(null);

  // Initialize auth state from localStorage
  useEffect(() => {
    const initializeAuth = () => {
      try {
        const stored = localStorage.getItem(TOKEN_STORAGE_KEY);
        if (stored) {
          const tokens: StoredTokens = JSON.parse(stored);
          // Check if token hasn't expired
          if (tokens.expiresAt > Date.now()) {
            setAccessToken(tokens.accessToken);
            setRefreshToken(tokens.refreshToken);
            setIsAuthenticated(true);
          } else {
            // Token expired, clear storage
            localStorage.removeItem(TOKEN_STORAGE_KEY);
            localStorage.removeItem(REFRESH_TOKEN_KEY);
          }
        }
      } catch (err) {
        console.error('Failed to initialize auth:', err);
        localStorage.removeItem(TOKEN_STORAGE_KEY);
        localStorage.removeItem(REFRESH_TOKEN_KEY);
      } finally {
        setIsLoading(false);
      }
    };

    initializeAuth();
  }, []);

  const storeTokens = useCallback((access: string, refresh: string) => {
    // Store tokens with 1 hour expiry
    const expiresAt = Date.now() + 60 * 60 * 1000;
    const tokens: StoredTokens = {
      accessToken: access,
      refreshToken: refresh,
      expiresAt,
    };
    localStorage.setItem(TOKEN_STORAGE_KEY, JSON.stringify(tokens));
    localStorage.setItem(REFRESH_TOKEN_KEY, refresh);
    setAccessToken(access);
    setRefreshToken(refresh);
    setIsAuthenticated(true);
  }, []);

  const clearTokens = useCallback(() => {
    localStorage.removeItem(TOKEN_STORAGE_KEY);
    localStorage.removeItem(REFRESH_TOKEN_KEY);
    setAccessToken(null);
    setRefreshToken(null);
    setIsAuthenticated(false);
  }, []);

  const signup = useCallback(
    async (email: string, password: string, userType: string, phoneNumber?: string) => {
      setIsLoading(true);
      setError(null);
      try {
        const response: SignupResponse = await apiClient.signup({
          email,
          password,
          user_type: userType,
          phone_number: phoneNumber,
        });

        if (response.error) {
          throw new Error(response.error.message);
        }

        // Signup now returns tokens directly!
        if (response.access_token && response.refresh_token) {
          storeTokens(response.access_token, response.refresh_token);
        } else {
          throw new Error('No tokens received from signup');
        }
      } catch (err) {
        const errorMessage = err instanceof Error ? err.message : 'Signup failed';
        setError(errorMessage);
        throw err;
      } finally {
        setIsLoading(false);
      }
    },
    [storeTokens]
  );

  const login = useCallback(async (email: string, password: string) => {
    setIsLoading(true);
    setError(null);
    try {
      const response: LoginResponse = await apiClient.login({
        email,
        password,
      });

      if (response.error) {
        throw new Error(response.error.message);
      }

      storeTokens(response.access_token, response.refresh_token);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Login failed';
      setError(errorMessage);
      throw err;
    } finally {
      setIsLoading(false);
    }
  }, [storeTokens]);

  const logout = useCallback(async () => {
    setIsLoading(true);
    setError(null);
    try {
      if (refreshToken) {
        await apiClient.logout(refreshToken);
      }
    } catch (err) {
      console.error('Logout API error:', err);
      // Still clear local tokens even if API call fails
    } finally {
      clearTokens();
      setIsLoading(false);
    }
  }, [refreshToken, clearTokens]);

  const clearError = useCallback(() => {
    setError(null);
  }, []);

  const value: AuthContextType = {
    isAuthenticated,
    isLoading,
    error,
    accessToken,
    refreshToken,
    signup,
    login,
    logout,
    clearError,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}
