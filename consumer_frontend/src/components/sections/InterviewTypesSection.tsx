'use client';

import { motion } from 'framer-motion';
import { Section } from '@/components/common/Section';
import { staggerContainer, staggerItem } from '@/utils/animations';
import { INTERVIEW_TYPES } from '@/constants/content';

export function InterviewTypesSection() {
    return (
        <Section className="bg-gradient-to-b from-gray-50 to-white">
            <div className="max-w-7xl mx-auto">
                <motion.div
                    className="text-center mb-16"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={staggerContainer}
                >
                    <span className="inline-block px-4 py-2 bg-orange-100 text-orange-600 rounded-full text-sm font-semibold mb-4">
                        📋 Interview Types
                    </span>
                    <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
                        Multiple Interview Formats
                    </h2>
                    <p className="text-xl text-gray-600 max-w-2xl mx-auto">
                        Practice all types of interviews in one platform
                    </p>
                </motion.div>

                <motion.div
                    className="grid grid-cols-1 md:grid-cols-3 gap-8"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={staggerContainer}
                >
                    {INTERVIEW_TYPES.map((type, index) => (
                        <motion.div
                            key={index}
                            variants={staggerItem}
                            whileHover={{ scale: 1.05, y: -10 }}
                            className="p-8 rounded-xl bg-white border border-gray-200 hover:border-blue-400 hover:shadow-xl transition-all"
                        >
                            <div className="text-5xl mb-4">
                                {index === 0 ? '💬' : index === 1 ? '🧠' : '💻'}
                            </div>
                            <h3 className="text-2xl font-bold text-gray-900 mb-4">{type.name}</h3>
                            <p className="text-gray-600 leading-relaxed">{type.description}</p>
                        </motion.div>
                    ))}
                </motion.div>
            </div>
        </Section>
    );
}
