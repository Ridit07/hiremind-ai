'use client';

import { motion } from 'framer-motion';
import { Button } from '@/components/common/Button';
import { fadeInUp } from '@/utils/animations';
import { HiArrowRight } from 'react-icons/hi';

export function CTASection() {
    return (
        <section className="py-20 px-4 md:px-8 lg:px-16 bg-gradient-to-r from-blue-600 to-orange-500">
            <motion.div
                className="max-w-4xl mx-auto text-center"
                initial="hidden"
                whileInView="visible"
                viewport={{ once: true }}
                variants={fadeInUp}
            >
                <h2 className="text-4xl md:text-5xl font-bold text-white mb-6">
                    Ready to Master Your Interviews?
                </h2>
                <p className="text-xl text-white/90 mb-8 max-w-2xl mx-auto">
                    Join thousands of successful candidates. Start your free trial today and transform your interview preparation.
                </p>
                <Button
                    href="/get-started"
                    variant="secondary"
                    size="lg"
                    className="!text-blue-600"
                >
                    Get Started for Free <HiArrowRight />
                </Button>
            </motion.div>
        </section>
    );
}
