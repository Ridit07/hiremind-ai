'use client';

import React, { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { motion } from 'framer-motion';
import { useAuth } from '@/hooks/useAuth';
import { useInterviews } from '@/hooks/useInterviews';
import { fadeInUp, staggerContainer, staggerItem } from '@/utils/animations';
import { HiCalendar, HiUser, HiCheckCircle, HiClock, HiXCircle } from 'react-icons/hi';

export default function InterviewsPage() {
    const router = useRouter();
    const { isAuthenticated, isLoading: authLoading } = useAuth();
    const { interviews, isLoading, error, refetch, clearError } = useInterviews();

    useEffect(() => {
        if (!authLoading && !isAuthenticated) {
            router.push('/login');
        }
    }, [isAuthenticated, authLoading, router]);

    if (authLoading || isLoading) {
        return (
            <main className="min-h-screen pt-20 pb-20 bg-gradient-to-br from-blue-50 to-white">
                <div className="max-w-6xl mx-auto px-4">
                    <motion.div
                        className="flex items-center justify-center h-96"
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                    >
                        <div className="text-center">
                            <div className="w-12 h-12 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
                            <p className="text-gray-600">Loading interviews...</p>
                        </div>
                    </motion.div>
                </div>
            </main>
        );
    }

    if (!isAuthenticated) {
        return null;
    }

    const getStatusIcon = (status: string) => {
        switch (status) {
            case 'scheduled':
                return <HiClock className="text-blue-600" />;
            case 'ongoing':
                return <HiClock className="text-orange-600" />;
            case 'completed':
                return <HiCheckCircle className="text-green-600" />;
            case 'cancelled':
                return <HiXCircle className="text-red-600" />;
            default:
                return <HiClock className="text-gray-600" />;
        }
    };

    const getStatusColor = (status: string) => {
        switch (status) {
            case 'scheduled':
                return 'bg-blue-50 text-blue-700 border-blue-200';
            case 'ongoing':
                return 'bg-orange-50 text-orange-700 border-orange-200';
            case 'completed':
                return 'bg-green-50 text-green-700 border-green-200';
            case 'cancelled':
                return 'bg-red-50 text-red-700 border-red-200';
            default:
                return 'bg-gray-50 text-gray-700 border-gray-200';
        }
    };

    const formatDate = (dateString: string) => {
        try {
            const date = new Date(dateString);
            return date.toLocaleDateString('en-US', {
                year: 'numeric',
                month: 'short',
                day: 'numeric',
                hour: '2-digit',
                minute: '2-digit',
            });
        } catch {
            return dateString;
        }
    };

    return (
        <main className="min-h-screen pt-20 pb-20 bg-gradient-to-br from-blue-50 to-white">
            <div className="max-w-6xl mx-auto px-4">
                {/* Header */}
                <motion.div
                    className="mb-12"
                    initial="hidden"
                    animate="visible"
                    variants={fadeInUp}
                >
                    <h1 className="text-4xl font-bold text-gray-900 mb-2">Your Interviews</h1>
                    <p className="text-gray-600">
                        {interviews.length} interview{interviews.length !== 1 ? 's' : ''} found
                    </p>
                </motion.div>

                {/* Error Message */}
                {error && (
                    <motion.div
                        className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg flex justify-between items-center"
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                    >
                        <p className="text-red-700">{error}</p>
                        <button
                            onClick={clearError}
                            className="text-red-600 hover:text-red-800 font-medium"
                        >
                            Dismiss
                        </button>
                    </motion.div>
                )}

                {/* Refresh Button */}
                <motion.div
                    className="mb-8"
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ delay: 0.1 }}
                >
                    <button
                        onClick={refetch}
                        disabled={isLoading}
                        className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed font-medium"
                    >
                        {isLoading ? 'Refreshing...' : 'Refresh'}
                    </button>
                </motion.div>

                {/* Interviews Grid */}
                {interviews.length === 0 ? (
                    <motion.div
                        className="bg-white rounded-xl border-2 border-dashed border-gray-300 p-12 text-center"
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                    >
                        <HiCalendar className="text-6xl text-gray-400 mx-auto mb-4" />
                        <h2 className="text-2xl font-bold text-gray-900 mb-2">No interviews yet</h2>
                        <p className="text-gray-600 mb-6">
                            You don't have any scheduled interviews at the moment.
                        </p>
                    </motion.div>
                ) : (
                    <motion.div
                        className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6"
                        initial="hidden"
                        animate="visible"
                        variants={staggerContainer}
                    >
                        {interviews.map((interview, index) => (
                            <motion.div
                                key={interview.interview_id}
                                variants={staggerItem}
                                className="bg-white rounded-xl shadow-md border border-gray-200 hover:shadow-lg transition-all overflow-hidden"
                                whileHover={{ y: -5 }}
                            >
                                {/* Header */}
                                <div className="bg-gradient-to-r from-blue-600 to-orange-500 px-6 py-4 text-white">
                                    <div className="flex items-center justify-between mb-2">
                                        <h3 className="text-lg font-bold">Interview</h3>
                                        <span className={`px-3 py-1 rounded-full text-sm font-medium border flex items-center gap-1 ${getStatusColor(interview.status)}`}>
                                            {getStatusIcon(interview.status)}
                                            {interview.status.charAt(0).toUpperCase() + interview.status.slice(1)}
                                        </span>
                                    </div>
                                </div>

                                {/* Content */}
                                <div className="px-6 py-4 space-y-4">
                                    {/* Candidate */}
                                    <div>
                                        <p className="text-xs text-gray-500 font-semibold uppercase tracking-wide mb-1">
                                            Candidate
                                        </p>
                                        <div className="flex items-center gap-2">
                                            <HiUser className="text-gray-400" />
                                            <p className="text-gray-900 font-medium">{interview.candidate_name}</p>
                                        </div>
                                    </div>

                                    {/* HR */}
                                    <div>
                                        <p className="text-xs text-gray-500 font-semibold uppercase tracking-wide mb-1">
                                            HR Professional
                                        </p>
                                        <p className="text-gray-700">{interview.hr_name}</p>
                                    </div>

                                    {/* Date & Time */}
                                    <div>
                                        <p className="text-xs text-gray-500 font-semibold uppercase tracking-wide mb-1">
                                            Interview Date & Time
                                        </p>
                                        <div className="flex items-center gap-2">
                                            <HiCalendar className="text-gray-400" />
                                            <p className="text-gray-700 text-sm">{formatDate(interview.interview_datetime)}</p>
                                        </div>
                                    </div>

                                    {/* Report */}
                                    {interview.interview_report_path && (
                                        <div>
                                            <p className="text-xs text-gray-500 font-semibold uppercase tracking-wide mb-1">
                                                Report
                                            </p>
                                            <a
                                                href={interview.interview_report_path}
                                                target="_blank"
                                                rel="noopener noreferrer"
                                                className="text-blue-600 hover:text-blue-700 text-sm font-medium underline"
                                            >
                                                View Report
                                            </a>
                                        </div>
                                    )}
                                </div>

                                {/* Footer */}
                                <div className="bg-gray-50 px-6 py-3 border-t border-gray-200 text-xs text-gray-500">
                                    Updated {formatDate(interview.updated_at)}
                                </div>
                            </motion.div>
                        ))}
                    </motion.div>
                )}
            </div>
        </main>
    );
}
