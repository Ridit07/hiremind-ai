'use client';

import { motion } from 'framer-motion';
import { FeatureCard } from '@/components/common/FeatureCard';
import { Section } from '@/components/common/Section';
import { staggerContainer, staggerItem } from '@/utils/animations';
import { FEATURES } from '@/constants/content';

export function FeaturesSection() {
    return (
        <Section className="bg-white">
            <div className="max-w-7xl mx-auto">
                <motion.div
                    className="text-center mb-16"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={staggerContainer}
                >
                    <span className="inline-block px-4 py-2 bg-blue-100 text-blue-600 rounded-full text-sm font-semibold mb-4">
                        � Why Companies Choose HireMind AI
                    </span>
                    <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
                        Streamline Your Entire Hiring Process
                    </h2>
                    <p className="text-xl text-gray-600 max-w-2xl mx-auto">
                        From candidate upload to final evaluation, everything you need to hire smarter
                    </p>
                </motion.div>

                <motion.div
                    className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8"
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    variants={staggerContainer}
                >
                    {FEATURES.map((feature, index) => (
                        <motion.div key={index} variants={staggerItem}>
                            <FeatureCard
                                icon={feature.icon}
                                title={feature.title}
                                description={feature.description}
                            />
                        </motion.div>
                    ))}
                </motion.div>
            </div>
        </Section>
    );
}
