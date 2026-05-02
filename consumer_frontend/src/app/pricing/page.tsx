'use client';

import { motion } from 'framer-motion';
import { Button } from '@/components/common/Button';
import { staggerContainer, staggerItem, fadeInUp } from '@/utils/animations';
import { HiCheckCircle } from 'react-icons/hi';

export default function Pricing() {
    const plans = [
        {
            name: 'Starter',
            price: 'Free',
            period: '',
            description: 'Perfect for getting started',
            features: [
                '5 interview sessions per month',
                'Behavioral interviews only',
                'Basic feedback',
                'Community support',
                'Progress tracking',
            ],
            cta: 'Start Free',
            highlighted: false,
        },
        {
            name: 'Professional',
            price: '$29',
            period: '/month',
            description: 'For serious candidates',
            features: [
                'Unlimited interview sessions',
                'All interview types',
                'Real-time AI feedback',
                'Company-specific practice',
                'Performance analytics',
                'Priority support',
                'Code evaluation',
                'Personalized recommendations',
            ],
            cta: 'Start 14-Day Free Trial',
            highlighted: true,
        },
        {
            name: 'Enterprise',
            price: 'Custom',
            period: '',
            description: 'For teams and organizations',
            features: [
                'Everything in Professional',
                'Team management',
                'Custom interview sets',
                'Advanced analytics',
                'API access',
                'Dedicated support',
                'White-label options',
                'Custom integrations',
            ],
            cta: 'Contact Sales',
            highlighted: false,
        },
    ];

    return (
        <main className="pt-20 pb-20">
            {/* Header */}
            <motion.section
                className="bg-gradient-to-br from-blue-50 to-white py-20 px-4"
                initial="hidden"
                animate="visible"
                variants={fadeInUp}
            >
                <div className="max-w-5xl mx-auto text-center">
                    <h1 className="text-5xl md:text-6xl font-bold text-gray-900 mb-6">
                        Simple, Transparent <span className="text-blue-600">Pricing</span>
                    </h1>
                    <p className="text-xl text-gray-600 max-w-2xl mx-auto mb-8">
                        Choose the perfect plan for your interview preparation journey
                    </p>

                    {/* Billing Toggle */}
                    <div className="inline-flex gap-4 bg-gray-100 rounded-full p-1">
                        <button className="px-6 py-2 rounded-full bg-white text-gray-900 font-semibold shadow-sm">
                            Monthly
                        </button>
                        <button className="px-6 py-2 rounded-full text-gray-600 font-semibold hover:text-gray-900">
                            Yearly <span className="text-green-600 ml-2">Save 20%</span>
                        </button>
                    </div>
                </div>
            </motion.section>

            {/* Pricing Cards */}
            <section className="py-20 px-4">
                <motion.div
                    className="max-w-6xl mx-auto"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={staggerContainer}
                >
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
                        {plans.map((plan, index) => (
                            <motion.div
                                key={index}
                                variants={staggerItem}
                                whileHover={{ y: -10 }}
                                className={`rounded-2xl p-8 border-2 transition-all ${plan.highlighted
                                        ? 'border-blue-600 bg-gradient-to-br from-blue-50 to-white shadow-2xl scale-105'
                                        : 'border-gray-200 bg-white hover:border-blue-300'
                                    }`}
                            >
                                {plan.highlighted && (
                                    <div className="mb-4 inline-block px-4 py-1 bg-blue-600 text-white rounded-full text-sm font-bold">
                                        Most Popular
                                    </div>
                                )}

                                <h3 className="text-2xl font-bold text-gray-900 mb-2">{plan.name}</h3>
                                <p className="text-gray-600 mb-6">{plan.description}</p>

                                <div className="mb-6">
                                    <span className="text-5xl font-bold text-gray-900">{plan.price}</span>
                                    {plan.period && <span className="text-gray-600">{plan.period}</span>}
                                </div>

                                <Button
                                    href="/get-started"
                                    variant={plan.highlighted ? 'primary' : 'outline'}
                                    className="w-full mb-8"
                                >
                                    {plan.cta}
                                </Button>

                                <div className="space-y-4">
                                    {plan.features.map((feature, idx) => (
                                        <motion.div
                                            key={idx}
                                            initial={{ opacity: 0, x: -10 }}
                                            whileInView={{ opacity: 1, x: 0 }}
                                            transition={{ delay: idx * 0.1 }}
                                            className="flex items-center gap-3"
                                        >
                                            <HiCheckCircle className="text-green-500 text-xl flex-shrink-0" />
                                            <span className="text-gray-700">{feature}</span>
                                        </motion.div>
                                    ))}
                                </div>
                            </motion.div>
                        ))}
                    </div>
                </motion.div>
            </section>

            {/* FAQ */}
            <section className="bg-gray-50 py-20 px-4">
                <motion.div
                    className="max-w-3xl mx-auto"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={fadeInUp}
                >
                    <h2 className="text-4xl font-bold text-gray-900 mb-12 text-center">Pricing FAQs</h2>

                    <div className="space-y-6">
                        {[
                            {
                                q: 'Can I upgrade or downgrade anytime?',
                                a: 'Yes, you can change your plan at any time. Changes take effect immediately.',
                            },
                            {
                                q: 'Is there a long-term contract?',
                                a: 'No contracts required. You can cancel your subscription anytime with no penalties.',
                            },
                            {
                                q: 'Do you offer refunds?',
                                a: 'We offer a 7-day money-back guarantee if you\'re not satisfied with our service.',
                            },
                            {
                                q: 'What payment methods do you accept?',
                                a: 'We accept all major credit cards, PayPal, and bank transfers for enterprise plans.',
                            },
                        ].map((faq, idx) => (
                            <motion.div
                                key={idx}
                                variants={staggerItem}
                                className="bg-white rounded-lg p-6 border border-gray-200"
                            >
                                <h3 className="text-lg font-bold text-gray-900 mb-2">{faq.q}</h3>
                                <p className="text-gray-600">{faq.a}</p>
                            </motion.div>
                        ))}
                    </div>
                </motion.div>
            </section>
        </main>
    );
}
