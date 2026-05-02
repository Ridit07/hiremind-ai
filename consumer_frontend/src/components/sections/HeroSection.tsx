'use client';

import { motion } from 'framer-motion';
import { Button } from '@/components/common/Button';
import { fadeInUp, slideInRight } from '@/utils/animations';
import { HiArrowRight, HiSparkles } from 'react-icons/hi';

export function HeroSection() {
    return (
        <section className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-purple-50 pt-32 pb-20 px-4 md:px-8 lg:px-16 overflow-hidden">
            <div className="max-w-7xl mx-auto">
                <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
                    {/* Left Content */}
                    <motion.div
                        initial="hidden"
                        animate="visible"
                        variants={fadeInUp}
                        className="space-y-6"
                    >
                        <motion.div
                            initial={{ opacity: 0, scale: 0.8 }}
                            animate={{ opacity: 1, scale: 1 }}
                            transition={{ duration: 0.6 }}
                        >
                            <span className="inline-block px-4 py-2 bg-blue-100 text-blue-600 rounded-full text-sm font-semibold mb-4 flex items-center gap-2">
                                <HiSparkles /> AI-Powered HR Solution
                            </span>
                        </motion.div>

                        <h1 className="text-5xl md:text-6xl lg:text-7xl font-bold text-gray-900 leading-tight">
                            Eliminate Interview <span className="text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-purple-600">
                                Hassles
                            </span>
                        </h1>

                        <p className="text-xl text-gray-600 max-w-xl">
                            Upload candidate lists, conduct AI-powered interviews, get detailed evaluations, and make better hiring decisions in minutes. The future of recruitment is here.
                        </p>

                        <div className="flex flex-col sm:flex-row gap-4 pt-8">
                            <Button href="/get-started" variant="primary" size="lg">
                                Start Free Trial <HiArrowRight />
                            </Button>
                            <Button href="/what-we-do" variant="outline" size="lg">
                                See How It Works
                            </Button>
                        </div>

                        <motion.div
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            transition={{ delay: 0.8, duration: 0.6 }}
                            className="flex gap-8 pt-8 border-t border-gray-200"
                        >
                            <div>
                                <p className="text-3xl font-bold text-blue-600">500+</p>
                                <p className="text-gray-600">Companies Hiring</p>
                            </div>
                            <div>
                                <p className="text-3xl font-bold text-purple-600">50K+</p>
                                <p className="text-gray-600">Candidates Evaluated</p>
                            </div>
                            <div>
                                <p className="text-3xl font-bold text-blue-600">85%</p>
                                <p className="text-gray-600">Time Saved</p>
                            </div>
                        </motion.div>
                    </motion.div>

                    {/* Right Illustration */}
                    <motion.div
                        initial="hidden"
                        animate="visible"
                        variants={slideInRight}
                        className="relative h-96 md:h-full"
                    >
                        <motion.div
                            className="absolute inset-0 bg-gradient-to-tr from-blue-400 to-purple-400 rounded-2xl blur-3xl opacity-20"
                            animate={{
                                scale: [1, 1.1, 1],
                                rotate: [0, 5, 0],
                            }}
                            transition={{ duration: 6, repeat: Infinity }}
                        />
                        <motion.div
                            className="relative bg-gradient-to-br from-blue-600 to-purple-600 rounded-2xl h-full flex items-center justify-center text-white text-6xl"
                            whileHover={{ scale: 1.05 }}
                            transition={{ duration: 0.3 }}
                        >
                            👥
                        </motion.div>
                    </motion.div>
                </div>
            </div>
        </section>
    );
}
