'use client';

import { motion } from 'framer-motion';
import { Section } from '@/components/common/Section';
import { staggerContainer, staggerItem } from '@/utils/animations';
import { TESTIMONIALS } from '@/constants/content';

export function TestimonialsSection() {
    return (
        <Section className="bg-gray-900">
            <div className="max-w-7xl mx-auto">
                <motion.div
                    className="text-center mb-16"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={staggerContainer}
                >
                    <span className="inline-block px-4 py-2 bg-blue-600/20 text-blue-400 rounded-full text-sm font-semibold mb-4">
                        ⭐ Success Stories
                    </span>
                    <h2 className="text-4xl md:text-5xl font-bold text-white mb-6">
                        What Users Say About Us
                    </h2>
                    <p className="text-xl text-gray-400 max-w-2xl mx-auto">
                        Join thousands of users who landed their dream jobs with HireMind AI
                    </p>
                </motion.div>

                <motion.div
                    className="grid grid-cols-1 md:grid-cols-3 gap-8"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={staggerContainer}
                >
                    {TESTIMONIALS.map((testimonial, index) => (
                        <motion.div
                            key={index}
                            variants={staggerItem}
                            whileHover={{ y: -10 }}
                            className="p-8 rounded-xl bg-gray-800 border border-gray-700 hover:border-blue-500 transition-all"
                        >
                            <div className="flex gap-1 mb-4">
                                {[...Array(5)].map((_, i) => (
                                    <span key={i} className="text-yellow-400">
                                        ★
                                    </span>
                                ))}
                            </div>
                            <p className="text-gray-300 mb-6 leading-relaxed">"{testimonial.quote}"</p>
                            <div className="flex items-center gap-4">
                                <div className="text-4xl">{testimonial.image}</div>
                                <div>
                                    <p className="font-bold text-white">{testimonial.author}</p>
                                    <p className="text-sm text-gray-400">{testimonial.role}</p>
                                </div>
                            </div>
                        </motion.div>
                    ))}
                </motion.div>
            </div>
        </Section>
    );
}
