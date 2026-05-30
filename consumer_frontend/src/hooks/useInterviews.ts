'use client';

import { useCallback, useState, useEffect } from 'react';
import { apiClient, InterviewData, GetInterviewsResponse } from '@/services/api';
import { useAuth } from './useAuth';

interface UseInterviewsReturn {
  interviews: InterviewData[];
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
  clearError: () => void;
}

export function useInterviews(): UseInterviewsReturn {
  const { accessToken, isAuthenticated } = useAuth();
  const [interviews, setInterviews] = useState<InterviewData[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchInterviews = useCallback(async () => {
    if (!accessToken || !isAuthenticated) {
      setInterviews([]);
      return;
    }

    setIsLoading(true);
    setError(null);

    try {
      const response: GetInterviewsResponse = await apiClient.getInterviews(accessToken);
      setInterviews(response.interviews || []);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch interviews';
      setError(errorMessage);
      setInterviews([]);
    } finally {
      setIsLoading(false);
    }
  }, [accessToken, isAuthenticated]);

  // Auto-fetch when authenticated
  useEffect(() => {
    if (isAuthenticated && accessToken) {
      fetchInterviews();
    }
  }, [isAuthenticated, accessToken, fetchInterviews]);

  const clearError = useCallback(() => {
    setError(null);
  }, []);

  return {
    interviews,
    isLoading,
    error,
    refetch: fetchInterviews,
    clearError,
  };
}
