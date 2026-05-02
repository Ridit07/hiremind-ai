'use client';

import { motion } from 'framer-motion';
import { Section } from '@/components/common/Section';
import { fadeInUp, staggerContainer, staggerItem } from '@/utils/animations';

export default function About() {
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
                        About <span className="text-blue-600">HireMind AI</span>
                    </h1>
                    <p className="text-xl text-gray-600 max-w-2xl mx-auto">
                        We're reimagining recruitment by using artificial intelligence to help companies hire the right talent faster and fairer.
                    </p>
                </motion.div>
            </Section>

            {/* Our Mission */}
            <Section className="bg-white">
                <div className="max-w-5xl mx-auto">
                    <motion.div
                        className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        <motion.div variants={staggerItem}>
                            <h2 className="text-4xl font-bold text-gray-900 mb-6">Our Mission</h2>
                            <p className="text-lg text-gray-600 mb-4">
                                We're on a mission to transform the hiring process. For too long, companies have relied on time-consuming, subjective interview processes that leave talent on the table.
                            </p>
                            <p className="text-lg text-gray-600 mb-4">
                                HireMind AI brings efficiency, fairness, and objectivity to recruitment. We help companies evaluate candidates at scale without sacrificing quality.
                            </p>
                            <p className="text-lg text-gray-600">
                                By leveraging advanced AI and proven assessment methodologies, we enable HR teams to make better hiring decisions in a fraction of the time.
                            </p>
                        </motion.div>
                        <motion.div
                            variants={staggerItem}
                            className="h-80 bg-gradient-to-br from-blue-400 to-purple-400 rounded-2xl flex items-center justify-center text-6xl"
                        >
                            🎯
                        </motion.div>
                    </motion.div>
                </div>
            </Section>

            {/* The Problem & Solution */}
            <Section className="bg-gray-50">
                <div className="max-w-5xl mx-auto">
                    <motion.div
                        className="text-center mb-16"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        <h2 className="text-4xl font-bold text-gray-900 mb-6">The Problem We Solve</h2>
                    </motion.div>

                    <motion.div
                        className="grid grid-cols-1 md:grid-cols-2 gap-8"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        {[
                            {
                                title: '❌ Manual Process',
                                description: 'HR teams spend weeks conducting interviews, leaving candidates waiting and good talent slipping away.',
                            },
                            {
                                title: '❌ Inconsistent Evaluations',
                                description: 'Subjective assessments lead to bias and unfair hiring decisions that don\'t reflect actual capability.',
                            },
                            {
                                title: '❌ High Costs',
                                description: 'Recruiting services, interviewer time, and extended hiring cycles drain your budget.',
                            },
                            {
                                title: '❌ Scalability Issues',
                                description: 'Hiring processes break down when you need to evaluate hundreds of candidates quickly.',
                            },
                        ].map((item, index) => (
                            <motion.div
                                key={index}
                                variants={staggerItem}
                                className="p-8 bg-white rounded-xl border border-gray-200"
                            >
                                <h3 className="text-xl font-bold text-gray-900 mb-3">{item.title}</h3>
                                <p className="text-gray-600">{item.description}</p>
                            </motion.div>
                        ))}
                    </motion.div>
                </div>
            </Section>

            {/* Why We're Different */}
            <Section className="bg-white">
                <div className="max-w-5xl mx-auto">
                    <motion.div
                        className="text-center mb-16"
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true }}
                        variants={staggerContainer}
                    >
                        <h2 className="text-4xl font-bold text-gray-900 mb-6">Why We're Different</h2>
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
                                icon: '⚡',
                                title: 'Speed',
                                description: 'Evaluate candidates in hours, not weeks. Scale your hiring without scaling your team.',
                            },
                            {
                                icon: '🎯',
                                title: 'Fairness',
                                description: 'Standardized evaluations eliminate bias. Every candidate is assessed by the same criteria.',
                            },
                            {
                                icon: '�',
                                title: 'Intelligence',
                                description: 'AI learns from your hiring patterns to suggest better candidates and predict success.',
                            },
                        ].map((value, index) => (
                            <motion.div
                                key={index}
                                variants={staggerItem}
                                className="p-8 bg-blue-50 rounded-xl border border-blue-200"
                            >
                                <div className="text-5xl mb-4">{value.icon}</div>
                                <h3 className="text-2xl font-bold text-gray-900 mb-3">{value.title}</h3>
                                <p className="text-gray-600">{value.description}</p>
                            </motion.div>
                        ))}
                    </motion.div>
                </div>
            </Section>
        </main>
    );
}
