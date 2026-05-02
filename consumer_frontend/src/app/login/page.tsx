'use client';

import { motion } from 'framer-motion';
import { fadeInUp } from '@/utils/animations';
import { HiMail, HiLockClosed } from 'react-icons/hi';

export default function Login() {
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
                        Welcome Back to <span className="text-blue-600">HireMind AI</span>
                    </h1>
                    <p className="text-gray-600">
                        Sign in to continue your interview preparation
                    </p>
                </motion.div>

                <motion.div
                    className="bg-white rounded-2xl shadow-xl border border-gray-200 p-8"
                    initial="hidden"
                    animate="visible"
                    variants={fadeInUp}
                    transition={{ delay: 0.2 }}
                >
                    {/* Login Form */}
                    <form className="space-y-6">
                        {/* Email */}
                        <motion.div
                            initial={{ opacity: 0, y: 10 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.3 }}
                        >
                            <label htmlFor="email" className="block text-sm font-semibold text-gray-900 mb-2">
                                Email Address
                            </label>
                            <div className="relative">
                                <HiMail className="absolute left-3 top-3 text-gray-400 text-xl" />
                                <input
                                    id="email"
                                    type="email"
                                    placeholder="you@example.com"
                                    className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent"
                                />
                            </div>
                        </motion.div>

                        {/* Password */}
                        <motion.div
                            initial={{ opacity: 0, y: 10 }}
                            animate={{ opacity: 1, y: 0 }}
                            transition={{ delay: 0.4 }}
                        >
                            <label htmlFor="password" className="block text-sm font-semibold text-gray-900 mb-2">
                                Password
                            </label>
                            <div className="relative">
                                <HiLockClosed className="absolute left-3 top-3 text-gray-400 text-xl" />
                                <input
                                    id="password"
                                    type="password"
                                    placeholder="••••••••"
                                    className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent"
                                />
                            </div>
                        </motion.div>

                        {/* Remember & Forgot */}
                        <motion.div
                            className="flex items-center justify-between text-sm"
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            transition={{ delay: 0.5 }}
                        >
                            <label className="flex items-center gap-2 cursor-pointer">
                                <input type="checkbox" className="w-4 h-4 rounded" />
                                <span className="text-gray-600">Remember me</span>
                            </label>
                            <a href="#" className="text-blue-600 hover:text-blue-700 font-medium">
                                Forgot Password?
                            </a>
                        </motion.div>

                        {/* Sign In Button */}
                        <motion.button
                            type="submit"
                            className="w-full bg-gradient-to-r from-blue-600 to-orange-500 text-white font-bold py-3 rounded-lg hover:shadow-lg transition-all duration-200"
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            transition={{ delay: 0.6 }}
                            whileHover={{ scale: 1.02 }}
                            whileTap={{ scale: 0.98 }}
                        >
                            Sign In
                        </motion.button>
                    </form>

                    {/* Divider */}
                    <motion.div
                        className="relative my-6"
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                        transition={{ delay: 0.7 }}
                    >
                        <div className="absolute inset-0 flex items-center">
                            <div className="w-full border-t border-gray-300"></div>
                        </div>
                        <div className="relative flex justify-center text-sm">
                            <span className="px-2 bg-white text-gray-500">Or continue with</span>
                        </div>
                    </motion.div>

                    {/* Social Login */}
                    <motion.div
                        className="grid grid-cols-3 gap-3"
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                        transition={{ delay: 0.8 }}
                    >
                        {['Google', 'GitHub', 'LinkedIn'].map((provider) => (
                            <button
                                key={provider}
                                className="py-2 px-3 border border-gray-300 rounded-lg hover:border-gray-400 hover:bg-gray-50 transition-all text-sm font-medium text-gray-700"
                            >
                                {provider === 'Google' && '🔍'}
                                {provider === 'GitHub' && '🐙'}
                                {provider === 'LinkedIn' && '💼'}
                            </button>
                        ))}
                    </motion.div>
                </motion.div>

                {/* Sign Up Link */}
                <motion.div
                    className="text-center mt-8"
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ delay: 0.9 }}
                >
                    <p className="text-gray-600">
                        Don't have an account?{' '}
                        <a href="/get-started" className="text-blue-600 hover:text-blue-700 font-bold">
                            Sign up for free
                        </a>
                    </p>
                </motion.div>

                {/* Security Notice */}
                <motion.div
                    className="mt-12 text-center text-sm text-gray-500"
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ delay: 1 }}
                >
                    <p>🔒 Your data is encrypted and secure. We never share your information.</p>
                </motion.div>
            </div>
        </main>
    );
}
