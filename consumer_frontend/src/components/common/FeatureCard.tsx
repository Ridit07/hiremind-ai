'use client';

import { motion } from 'framer-motion';

interface CardProps {
    icon?: string | React.ReactNode;
    title: string;
    description: string;
    className?: string;
}

export function FeatureCard({ icon, title, description, className = '' }: CardProps) {
    return (
        <motion.div
            className={`p-6 rounded-xl bg-white border border-gray-200 hover:border-blue-400 hover:shadow-lg transition-all duration-300 ${className}`}
            whileHover={{ y: -8 }}
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5 }}
            viewport={{ once: true }}
        >
            {icon && (
                <div className="text-4xl mb-4">
                    {typeof icon === 'string' ? icon : icon}
                </div>
            )}
            <h3 className="text-xl font-bold text-gray-900 mb-3">{title}</h3>
            <p className="text-gray-600">{description}</p>
        </motion.div>
    );
}
