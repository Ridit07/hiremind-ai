'use client';

import { motion } from 'framer-motion';
import { Button } from '@/components/common/Button';
import { fadeInUp, staggerContainer, staggerItem } from '@/utils/animations';

export default function GetStarted() {
    return (
        <main className="pt-20 pb-20 bg-gradient-to-br from-blue-50 to-white">
            <div className="max-w-5xl mx-auto px-4">
                {/* Header */}
                <motion.div
                    className="text-center mb-16"
                    initial="hidden"
                    animate="visible"
                    variants={fadeInUp}
                >
                    <h1 className="text-5xl md:text-6xl font-bold text-gray-900 mb-6">
                        Get Started with <span className="text-blue-600">HireMind AI</span>
                    </h1>
                    <p className="text-xl text-gray-600 max-w-2xl mx-auto">
                        Join 500+ companies transforming their hiring process
                    </p>
                </motion.div>

                {/* Steps */}
                <motion.div
                    className="space-y-8 mb-20"
                    initial="hidden"
                    animate="visible"
                    variants={staggerContainer}
                >
                    {[
                        {
                            step: 1,
                            title: 'Sign Up Your Team',
                            description: 'Create a company account and invite HR team members. Takes less than 5 minutes.',
                        },
                        {
                            step: 2,
                            title: 'Configure Your Interviews',
                            description: 'Set interview types, difficulty levels, duration, and add your custom questions.',
                        },
                        {
                            step: 3,
                            title: 'Upload Candidates',
                            description: 'Import candidate lists from your ATS or upload manually. Support for CSV, Excel, and API.',
                        },
                        {
                            step: 4,
                            title: 'Launch Interviews',
                            description: 'Send automated interviews to candidates. They complete on their own schedule.',
                        },
                        {
                            step: 5,
                            title: 'Review Reports',
                            description: 'Get comprehensive evaluation reports with scores, insights, and recommendations.',
                        },
                        {
                            step: 6,
                            title: 'Make Offers',
                            description: 'Use data-driven insights to make better hiring decisions and extend offers.',
                        },
                    ].map((item, index) => (
                        <motion.div
                            key={index}
                            variants={staggerItem}
                            whileHover={{ x: 10 }}
                            className="flex gap-8 items-start p-8 bg-white rounded-xl border border-gray-200 hover:border-blue-400 hover:shadow-lg transition-all"
                        >
                            <div className="flex-shrink-0">
                                <div className="w-16 h-16 rounded-full bg-gradient-to-br from-blue-600 to-orange-500 text-white flex items-center justify-center font-bold text-2xl">
                                    {item.step}
                                </div>
                            </div>
                            <div className="flex-1">
                                <h3 className="text-2xl font-bold text-gray-900 mb-3">{item.title}</h3>
                                <p className="text-lg text-gray-600">{item.description}</p>
                            </div>
                        </motion.div>
                    ))}
                </motion.div>

                {/* Benefits */}
                <motion.div
                    className="bg-gradient-to-r from-blue-600 to-purple-600 rounded-2xl p-12 text-white mb-12"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={fadeInUp}
                >
                    <h2 className="text-3xl font-bold mb-8">What You'll Get</h2>
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                        {[
                            { icon: '⚡', text: '70% faster hiring process' },
                            { icon: '📊', text: 'Standardized candidate evaluation' },
                            { icon: '💰', text: 'Reduced recruiting costs' },
                            { icon: '🎯', text: 'Better quality hires' },
                            { icon: '📈', text: 'Detailed candidate analytics' },
                            { icon: '🔗', text: 'ATS integration support' },
                            { icon: '👥', text: 'Unlimited candidate evaluations' },
                            { icon: '🔒', text: 'Enterprise-grade security' },
                        ].map((benefit, index) => (
                            <motion.div
                                key={index}
                                variants={staggerItem}
                                className="flex items-center gap-3"
                            >
                                <span className="text-2xl">{benefit.icon}</span>
                                <span className="text-lg">{benefit.text}</span>
                            </motion.div>
                        ))}
                    </div>
                </motion.div>

                {/* CTA */}
                <motion.div
                    className="text-center"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={fadeInUp}
                >
                    <Button href="/login" variant="primary" size="lg" className="mb-6">
                        Start Your Free Trial
                    </Button>
                    <p className="text-gray-600">
                        14-day free trial. No credit card required. Full access to all features.
                    </p>
                </motion.div>

                {/* FAQ Section */}
                <motion.div
                    className="mt-20 bg-white rounded-xl border border-gray-200 p-12"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={fadeInUp}
                >
                    <h2 className="text-3xl font-bold text-gray-900 mb-8">Frequently Asked Questions</h2>
                    <div className="space-y-6">
                        {[
                            {
                                q: 'How many candidates can we evaluate?',
                                a: 'Unlimited. Evaluate as many candidates as you need with our platform. No per-candidate fees.',
                            },
                            {
                                q: 'Can we integrate with our existing ATS?',
                                a: 'Yes! We integrate with popular ATS platforms including Workday, Taleo, Greenhouse, Lever, and more.',
                            },
                            {
                                q: 'How long does an interview take?',
                                a: 'Typically 30-60 minutes depending on the interview type. Candidates can take interviews on their own schedule.',
                            },
                            {
                                q: 'Can we customize the questions?',
                                a: 'Absolutely. Add your own questions, adjust difficulty levels, and tailor interviews to your specific needs.',
                            },
                        ].map((faq, index) => (
                            <motion.div
                                key={index}
                                variants={staggerItem}
                                className="pb-6 border-b border-gray-200 last:border-0"
                            >
                                <h3 className="text-xl font-bold text-gray-900 mb-3">{faq.q}</h3>
                                <p className="text-gray-600">{faq.a}</p>
                            </motion.div>
                        ))}
                    </div>
                </motion.div>
            </div>
        </main>
    );
}
