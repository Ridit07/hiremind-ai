'use client';

import { ReactNode } from 'react';
import { motion } from 'framer-motion';
import { cn } from '@/utils/cn';

interface ButtonProps {
    children: ReactNode;
    href?: string;
    onClick?: () => void;
    variant?: 'primary' | 'secondary' | 'outline';
    size?: 'sm' | 'md' | 'lg';
    className?: string;
    disabled?: boolean;
    type?: 'button' | 'submit' | 'reset';
}

export function Button({
    children,
    href,
    onClick,
    variant = 'primary',
    size = 'md',
    className,
    disabled = false,
    type = 'button',
}: ButtonProps) {
    const baseStyles =
        'font-semibold rounded-lg transition-all duration-200 inline-flex items-center justify-center gap-2 cursor-pointer';

    const variants = {
        primary: 'bg-blue-600 text-white hover:bg-blue-700 shadow-lg hover:shadow-xl',
        secondary: 'bg-white text-blue-600 hover:bg-gray-100 border-2 border-blue-600',
        outline: 'border-2 border-blue-600 text-blue-600 hover:bg-blue-50',
    };

    const sizes = {
        sm: 'px-4 py-2 text-sm',
        md: 'px-6 py-3 text-base',
        lg: 'px-8 py-4 text-lg',
    };

    const buttonClass = cn(baseStyles, variants[variant], sizes[size], className);

    if (href) {
        return (
            <motion.a
                href={href}
                className={buttonClass}
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
            >
                {children}
            </motion.a>
        );
    }

    return (
        <motion.button
            type={type}
            className={buttonClass}
            onClick={onClick}
            disabled={disabled}
            whileHover={{ scale: 1.05 }}
            whileTap={{ scale: 0.95 }}
        >
            {children}
        </motion.button>
    );
}
