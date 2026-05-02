'use client';

import { motion } from 'framer-motion';
import { Section } from '@/components/common/Section';
import { FeatureCard } from '@/components/common/FeatureCard';
import { Button } from '@/components/common/Button';
import { staggerContainer, staggerItem, fadeInUp } from '@/utils/animations';
import { HiArrowRight } from 'react-icons/hi';

export default function WhatWeDo() {
    return (
        <main className="pt-20">
            {/* Hero */}
            <Section className="bg-gradient-to-br from-blue-50 to-white">
                <motion.div
                    className="max-w-4xl mx-auto text-center"
                    initial="hidden"
                    animate="visible"
                    variants={fadeInUp}
                >
                    <h1 className="text-5xl md:text-6xl font-bold text-gray-900 mb-6">
                        How <span className="text-blue-600">HireMind AI</span> Works
                    </h1>
                    <p className="text-xl text-gray-600 max-w-2xl mx-auto">
                        A complete hiring solution that eliminates the guesswork from recruitment
                    </p>
                </motion.div>
            </Section>

            {/* The Process */}
            <Section className="bg-white">
                <div className="max-w-5xl mx-auto">
                    <motion.div
                        className="text-center mb-16"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        <h2 className="text-4xl font-bold text-gray-900 mb-6">The HireMind Process</h2>
                    </motion.div>

                    <motion.div
                        className="space-y-8"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        {[
                            {
                                step: 1,
                                title: 'Upload Candidates',
                                description: 'Import your candidate list from your ATS or manually upload CVs. We support CSV, Excel, and direct ATS integration.',
                            },
                            {
                                step: 2,
                                title: 'Configure Interview',
                                description: 'Set interview parameters: type (behavioral, technical, coding), duration, difficulty level, and company-specific questions.',
                            },
                            {
                                step: 3,
                                title: 'Run Automated Interviews',
                                description: 'AI interviewers conduct interviews with each candidate. Real-time interaction, adaptive questioning, and complete recording.',
                            },
                            {
                                step: 4,
                                title: 'Get Instant Reports',
                                description: 'Receive comprehensive evaluation reports with scores, insights, and recommendations for each candidate.',
                            },
                            {
                                step: 5,
                                title: 'Make Data-Driven Decisions',
                                description: 'Compare candidates side-by-side, view detailed analytics, and make hiring decisions with confidence.',
                            },
                        ].map((item, index) => (
                            <motion.div
                                key={index}
                                variants={staggerItem}
                                className="flex gap-6 items-start"
                            >
                                <div className="flex-shrink-0">
                                    <div className="w-12 h-12 rounded-full bg-blue-600 text-white flex items-center justify-center font-bold text-lg">
                                        {item.step}
                                    </div>
                                </div>
                                <div className="flex-1 pt-2">
                                    <h3 className="text-2xl font-bold text-gray-900 mb-2">{item.title}</h3>
                                    <p className="text-gray-600 text-lg">{item.description}</p>
                                </div>
                            </motion.div>
                        ))}
                    </motion.div>
                </div>
            </Section>

            {/* Interview Types */}
            <Section className="bg-gray-50">
                <div className="max-w-5xl mx-auto">
                    <motion.div
                        className="text-center mb-16"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        <h2 className="text-4xl font-bold text-gray-900 mb-6">Interview Types Supported</h2>
                    </motion.div>

                    <motion.div
                        className="grid grid-cols-1 md:grid-cols-3 gap-8"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        {[
                            {
                                icon: '💬',
                                title: 'Behavioral',
                                description: 'Assess soft skills, cultural fit, and communication through scenario-based questions.',
                            },
                            {
                                icon: '🧮',
                                title: 'Technical',
                                description: 'Evaluate technical knowledge, problem-solving, and system design thinking.',
                            },
                            {
                                icon: '💻',
                                title: 'Coding',
                                description: 'Test coding skills with live coding challenges and automated code evaluation.',
                            },
                        ].map((type, index) => (
                            <motion.div
                                key={index}
                                variants={staggerItem}
                                whileHover={{ scale: 1.05 }}
                                className="p-8 bg-white rounded-xl border border-gray-200"
                            >
                                <div className="text-5xl mb-4">{type.icon}</div>
                                <h3 className="text-2xl font-bold text-gray-900 mb-3">{type.title}</h3>
                                <p className="text-gray-600">{type.description}</p>
                            </motion.div>
                        ))}
                    </motion.div>
                </div>
            </Section>

            {/* Key Capabilities */}
            <Section className="bg-white">
                <div className="max-w-5xl mx-auto">
                    <motion.div
                        className="text-center mb-16"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        <h2 className="text-4xl font-bold text-gray-900 mb-6">Key Capabilities</h2>
                    </motion.div>

                    <motion.div
                        className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        {[
                            {
                                icon: '📤',
                                title: 'Bulk Import',
                                description: 'Import candidate lists from your ATS in seconds',
                            },
                            {
                                icon: '🤖',
                                title: 'AI Interviewers',
                                description: 'Sophisticated AI that conducts human-like interviews',
                            },
                            {
                                icon: '📊',
                                title: 'Analytics Dashboard',
                                description: 'Comprehensive reporting and candidate comparison',
                            },
                            {
                                icon: '🎯',
                                title: 'Custom Evaluations',
                                description: 'Tailor interviews to your specific role requirements',
                            },
                            {
                                icon: '�',
                                title: 'ATS Integration',
                                description: 'Seamless integration with popular ATS platforms',
                            },
                            {
                                icon: '⚖️',
                                title: 'Bias Reduction',
                                description: 'Standardized evaluations for fairer hiring',
                            },
                        ].map((capability, index) => (
                            <motion.div
                                key={index}
                                variants={staggerItem}
                            >
                                <FeatureCard
                                    icon={capability.icon}
                                    title={capability.title}
                                    description={capability.description}
                                />
                            </motion.div>
                        ))}
                    </motion.div>
                </div>
            </Section>

            {/* CTA */}
            <Section className="bg-gradient-to-r from-blue-600 to-purple-600">
                <motion.div
                    className="max-w-4xl mx-auto text-center"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={fadeInUp}
                >
                    <h2 className="text-4xl md:text-5xl font-bold text-white mb-6">
                        Ready to Transform Your Hiring?
                    </h2>
                    <Button
                        href="/get-started"
                        variant="secondary"
                        size="lg"
                        className="!text-blue-600"
                    >
                        Start Your Free Trial <HiArrowRight />
                    </Button>
                </motion.div>
            </Section>
        </main>
    );
}
