'use client';

import Link from 'next/link';
import { useState } from 'react';
import { motion } from 'framer-motion';
import { HiMenu, HiX } from 'react-icons/hi';
import { Button } from './Button';
import { NAV_LINKS, CTA_BUTTONS } from '@/constants/navigation';

export function Navbar() {
    const [isOpen, setIsOpen] = useState(false);
    return (
        <nav className="fixed top-0 w-full bg-white/95 backdrop-blur-md shadow-md z-50">
            <div className="max-w-7xl mx-auto px-4 md:px-8 lg:px-16">
                <div className="flex justify-between items-center h-16 md:h-20">
                    {/* Logo */}
                    <Link href="/" className="text-2xl font-bold text-blue-600">
                        HireMind<span className="text-orange-500">AI</span>
                    </Link>

                    {/* Desktop Navigation */}
                    <div className="hidden md:flex items-center gap-8">
                        {NAV_LINKS.map((link) => (
                            <Link
                                key={link.href}
                                href={link.href}
                                className="text-gray-700 hover:text-blue-600 transition-colors font-medium"
                            >
                                {link.label}
                            </Link>
                        ))}
                    </div>

                    {/* Desktop CTA Buttons */}
                    <div className="hidden md:flex items-center gap-4">
                        <Button href={CTA_BUTTONS.secondary.href} variant="outline" size="sm">
                            {CTA_BUTTONS.secondary.label}
                        </Button>
                        <Button href={CTA_BUTTONS.primary.href} variant="primary" size="sm">
                            {CTA_BUTTONS.primary.label}
                        </Button>
                    </div>

                    {/* Mobile Menu Button */}
                    <button
                        className="md:hidden text-gray-700 text-2xl"
                        onClick={() => setIsOpen(!isOpen)}
                    >
                        {isOpen ? <HiX /> : <HiMenu />}
                    </button>
                </div>

                {/* Mobile Navigation */}
                {isOpen && (
                    <motion.div
                        initial={{ opacity: 0, y: -10 }}
                        animate={{ opacity: 1, y: 0 }}
                        exit={{ opacity: 0, y: -10 }}
                        className="md:hidden pb-4 border-t border-gray-200"
                    >
                        {NAV_LINKS.map((link) => (
                            <Link
                                key={link.href}
                                href={link.href}
                                className="block py-2 text-gray-700 hover:text-blue-600 transition-colors"
                                onClick={() => setIsOpen(false)}
                            >
                                {link.label}
                            </Link>
                        ))}
                        <div className="flex gap-2 mt-4">
                            <Button href={CTA_BUTTONS.secondary.href} variant="outline" size="sm">
                                {CTA_BUTTONS.secondary.label}
                            </Button>
                            <Button href={CTA_BUTTONS.primary.href} variant="primary" size="sm">
                                {CTA_BUTTONS.primary.label}
                            </Button>
                        </div>
                    </motion.div>
                )}
            </div>
        </nav>
    );
}
