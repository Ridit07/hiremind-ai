
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

// Helper function to decode JWT and extract user_id
function extractUserIdFromToken(token: string): string {
    try {
        const parts = token.split('.');
        if (parts.length !== 3) {
            console.warn('Invalid token format');
            return '';
        }

        const base64 = parts[1].replace(/-/g, '+').replace(/_/g, '/');
        const padded = base64.padEnd(base64.length + (4 - (base64.length % 4)) % 4, '=');
        const decoded = JSON.parse(atob(padded));
        return decoded.sub || decoded.user_id || decoded.userId || '';
    } catch (error) {
        console.warn('Failed to decode token:', error);
        return '';
    }
}

export interface SignupRequest {
    email: string;
    password: string;
    user_type: string;
    phone_number?: string;
}

export interface SignupResponse {
    message: string;
    access_token?: string;
    refresh_token?: string;
    error?: {
        code: number;
        type: string;
        message: string;
    };
}

export interface LoginRequest {
    email: string;
    password: string;
}

export interface LoginResponse {
    access_token: string;
    refresh_token: string;
    error?: {
        code: number;
        type: string;
        message: string;
    };
}

export interface RefreshTokenRequest {
    refresh_token: string;
}

export interface RefreshTokenResponse {
    access_token: string;
    refresh_token: string;
    error?: {
        code: number;
        type: string;
        message: string;
    };
}

export interface LogoutRequest {
    refresh_token: string;
}

export interface LogoutResponse {
    message: string;
    error?: {
        code: number;
        type: string;
        message: string;
    };
}

export interface InterviewData {
    interview_id: string;
    hr_id: string;
    hr_name: string;
    candidate_id: string;
    candidate_name: string;
    status: string;
    interview_datetime: string;
    interview_report_path: string;
    created_at: string;
    updated_at: string;
}

export interface GetInterviewsResponse {
    interviews: InterviewData[];
}

class APIClient {
    private baseURL: string;

    constructor(baseURL: string = API_BASE_URL) {
        this.baseURL = baseURL;
    }

    private async request<T>(
        endpoint: string,
        options: RequestInit = {}
    ): Promise<T> {
        const url = `${this.baseURL}${endpoint}`;

        const defaultHeaders: HeadersInit = {
            'Content-Type': 'application/json',
        };

        const response = await fetch(url, {
            ...options,
            headers: {
                ...defaultHeaders,
                ...(options.headers || {}),
            },
        });

        if (!response.ok) {
            const error = await response.json().catch(() => ({
                message: response.statusText,
            }));
            throw new Error(error.message || `API error: ${response.status}`);
        }

        return response.json() as Promise<T>;
    }

    // Auth Endpoints
    async signup(payload: SignupRequest): Promise<SignupResponse> {
        return this.request<SignupResponse>('/auth/signup', {
            method: 'POST',
            body: JSON.stringify(payload),
        });
    }

    async login(payload: LoginRequest): Promise<LoginResponse> {
        return this.request<LoginResponse>('/auth/login', {
            method: 'POST',
            body: JSON.stringify(payload),
        });
    }

    async refreshToken(refreshToken: string): Promise<RefreshTokenResponse> {
        return this.request<RefreshTokenResponse>('/auth/refresh', {
            method: 'POST',
            body: JSON.stringify({ refresh_token: refreshToken }),
        });
    }

    async logout(refreshToken: string): Promise<LogoutResponse> {
        return this.request<LogoutResponse>('/auth/logout', {
            method: 'POST',
            body: JSON.stringify({ refresh_token: refreshToken }),
        });
    }

    // Interview Endpoints
    async getInterviews(accessToken: string): Promise<GetInterviewsResponse> {
        const userId = extractUserIdFromToken(accessToken);

        return this.request<GetInterviewsResponse>('/interviews', {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${accessToken}`,
                'X-User-ID': userId,
            },
        });
    }
}

// Export singleton instance
export const apiClient = new APIClient();
