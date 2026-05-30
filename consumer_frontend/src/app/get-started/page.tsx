'use client';

import React, { useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { motion } from 'framer-motion';
import { useAuth } from '@/hooks/useAuth';
import { fadeInUp } from '@/utils/animations';
import { HiMail, HiLockClosed, HiUser, HiPhone } from 'react-icons/hi';

export default function SignupPage() {
    const router = useRouter();
    const { signup, isLoading, error, clearError } = useAuth();
    const [formData, setFormData] = useState({
        email: '',
        password: '',
        confirmPassword: '',
        userType: 'CANDIDATE',
        phoneNumber: '',
    });
    const [localError, setLocalError] = useState<string | null>(null);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLocalError(null);
        clearError();

        // Validation
        if (!formData.email || !formData.password || !formData.confirmPassword) {
            setLocalError('Please fill in all required fields');
            return;
        }

        if (formData.password !== formData.confirmPassword) {
            setLocalError('Passwords do not match');
            return;
        }

        if (formData.password.length < 6) {
            setLocalError('Password must be at least 6 characters');
            return;
        }

        try {
            await signup(formData.email, formData.password, formData.userType, formData.phoneNumber);
            router.push('/');
        } catch (err) {
            const errorMessage = err instanceof Error ? err.message : 'Signup failed';
            setLocalError(errorMessage);
        }
    };

    const displayError = localError || error;

    return (
        <main className="min-h-screen pt-20 pb-20 bg-gradient-to-br from-blue-50 to-white">
            <div className="max-w-md w-full mx-auto px-4">
                <motion.div
                    className="text-center mb-12"
                    initial="hidden"
                    animate="visible"
                    variants={fadeInUp}
                >
                    <h1 className="text-4xl font-bold text-gray-900 mb-2">
                        Join <span className="text-blue-600">HireMind AI</span>
                    </h1>
                    <p className="text-gray-600">
                        Create your account to get started today
                    </p>
                </motion.div>

                <motion.div
                    className="bg-white rounded-2xl shadow-xl border border-gray-200 p-8"
                    initial="hidden"
                    animate="visible"
                    variants={fadeInUp}
                    transition={{ delay: 0.2 }}
                >
                    {displayError && (
                        <motion.div
                            className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg"
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                        >
                            <p className="text-red-700 text-sm">{displayError}</p>
                        </motion.div>
                    )}

                    <form onSubmit={handleSubmit} className="space-y-5">
                        {/* User Type */}
                        <motion.div
                            initial={{ opacity: 0, y: 10 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.3 }}
                        >
                            <label htmlFor="userType" className="block text-sm font-semibold text-gray-900 mb-2">
                                Account Type
                            </label>
                            <select
                                id="userType"
                                name="userType"
                                value={formData.userType}
                                onChange={handleChange}
                                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent disabled:bg-gray-100 disabled:cursor-not-allowed"
                                disabled={isLoading}
                            >
                                <option value="CANDIDATE">Candidate</option>
                                <option value="HR">HR Professional</option>
                            </select>
                        </motion.div>

                        {/* Email */}
                        <motion.div
                            initial={{ opacity: 0, y: 10 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.35 }}
                        >
                            <label htmlFor="email" className="block text-sm font-semibold text-gray-900 mb-2">
                                Email Address
                            </label>
                            <div className="relative">
                                <HiMail className="absolute left-3 top-3 text-gray-400 text-xl" />
                                <input
                                    id="email"
                                    type="email"
                                    name="email"
                                    value={formData.email}
                                    onChange={handleChange}
                                    placeholder="you@example.com"
                                    className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent disabled:bg-gray-100 disabled:cursor-not-allowed"
                                    disabled={isLoading}
                                />
                            </div>
                        </motion.div>

                        {/* Phone Number (Optional) */}
                        <motion.div
                            initial={{ opacity: 0, y: 10 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.4 }}
                        >
                            <label htmlFor="phoneNumber" className="block text-sm font-semibold text-gray-900 mb-2">
                                Phone Number (Optional)
                            </label>
                            <div className="relative">
                                <HiPhone className="absolute left-3 top-3 text-gray-400 text-xl" />
                                <input
                                    id="phoneNumber"
                                    type="tel"
                                    name="phoneNumber"
                                    value={formData.phoneNumber}
                                    onChange={handleChange}
                                    placeholder="+1 (555) 000-0000"
                                    className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent disabled:bg-gray-100 disabled:cursor-not-allowed"
                                    disabled={isLoading}
                                />
                            </div>
                        </motion.div>

                        {/* Password */}
                        <motion.div
                            initial={{ opacity: 0, y: 10 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.45 }}
                        >
                            <label htmlFor="password" className="block text-sm font-semibold text-gray-900 mb-2">
                                Password
                            </label>
                            <div className="relative">
                                <HiLockClosed className="absolute left-3 top-3 text-gray-400 text-xl" />
                                <input
                                    id="password"
                                    type="password"
                                    name="password"
                                    value={formData.password}
                                    onChange={handleChange}
                                    placeholder="••••••••"
                                    className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent disabled:bg-gray-100 disabled:cursor-not-allowed"
                                    disabled={isLoading}
                                />
                            </div>
                        </motion.div>

                        {/* Confirm Password */}
                        <motion.div
                            initial={{ opacity: 0, y: 10 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.5 }}
                        >
                            <label htmlFor="confirmPassword" className="block text-sm font-semibold text-gray-900 mb-2">
                                Confirm Password
                            </label>
                            <div className="relative">
                                <HiLockClosed className="absolute left-3 top-3 text-gray-400 text-xl" />
                                <input
                                    id="confirmPassword"
                                    type="password"
                                    name="confirmPassword"
                                    value={formData.confirmPassword}
                                    onChange={handleChange}
                                    placeholder="••••••••"
                                    className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent disabled:bg-gray-100 disabled:cursor-not-allowed"
                                    disabled={isLoading}
                                />
                            </div>
                        </motion.div>

                        {/* Terms Agreement */}
                        <motion.div
                            className="flex items-start gap-2"
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            transition={{ delay: 0.55 }}
                        >
                            <input
                                type="checkbox"
                                id="terms"
                                className="w-4 h-4 rounded mt-0.5 cursor-pointer"
                                required
                            />
                            <label htmlFor="terms" className="text-sm text-gray-600 cursor-pointer">
                                I agree to the{' '}
                                <Link href="/terms" className="text-blue-600 hover:text-blue-700 font-medium">
                                    Terms of Service
                                </Link>{' '}
                                and{' '}
                                <Link href="/privacy" className="text-blue-600 hover:text-blue-700 font-medium">
                                    Privacy Policy
                                </Link>
                            </label>
                        </motion.div>

                        {/* Sign Up Button */}
                        <motion.button
                            type="submit"
                            disabled={isLoading}
                            className="w-full bg-gradient-to-r from-blue-600 to-orange-500 text-white font-bold py-3 rounded-lg hover:shadow-lg transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed mt-6"
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            transition={{ delay: 0.6 }}
                            whileHover={{ scale: isLoading ? 1 : 1.02 }}
                            whileTap={{ scale: isLoading ? 1 : 0.98 }}
                        >
                            {isLoading ? 'Creating Account...' : 'Create Account'}
                        </motion.button>
                    </form>
                </motion.div>

                {/* Login Link */}
                <motion.div
                    className="text-center mt-8"
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ delay: 0.7 }}
                >
                    <p className="text-gray-600">
                        Already have an account?{' '}
                        <Link href="/login" className="text-blue-600 hover:text-blue-700 font-bold">
                            Login here
                        </Link>
                    </p>
                </motion.div>

                {/* Security Notice */}
                <motion.div
                    className="mt-12 text-center text-sm text-gray-500"
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ delay: 0.8 }}
                >
                    <p>🔒 Your data is encrypted and secure. We never share your information.</p>
                </motion.div>
            </div>
        </main>
    );
}
