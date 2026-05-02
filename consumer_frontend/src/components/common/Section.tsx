'use client';

import { ReactNode } from 'react';
import { motion, Variants } from 'framer-motion';

interface SectionProps {
    children: ReactNode;
    className?: string;
    variants?: Variants;
    id?: string;
}

export function Section({ children, className = '', variants, id }: SectionProps) {
    return (
        <motion.section
            id={id}
            className={`py-20 px-4 md:px-8 lg:px-16 ${className}`}
            initial="hidden"
            whileInView="visible"
            viewport={{ once: true, margin: '-100px' }}
            variants={variants}
        >
            {children}
        </motion.section>
    );
}
