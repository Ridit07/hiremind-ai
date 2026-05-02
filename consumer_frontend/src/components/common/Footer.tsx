'use client';

import Link from 'next/link';
import { motion } from 'framer-motion';
import { HiArrowRight } from 'react-icons/hi';

export function Footer() {
    const currentYear = new Date().getFullYear();

    const footerLinks = {
        Product: [
            { label: 'Features', href: '#' },
            { label: 'Pricing', href: '/pricing' },
            { label: 'Security', href: '#' },
            { label: 'Roadmap', href: '#' },
        ],
        Company: [
            { label: 'About', href: '/about' },
            { label: 'Blog', href: '#' },
            { label: 'Careers', href: '#' },
            { label: 'Contact', href: '/contact' },
        ],
        Resources: [
            { label: 'Documentation', href: '#' },
            { label: 'API Reference', href: '#' },
            { label: 'Community', href: '#' },
            { label: 'Support', href: '#' },
        ],
        Legal: [
            { label: 'Privacy', href: '#' },
            { label: 'Terms', href: '#' },
            { label: 'Cookie Policy', href: '#' },
            { label: 'License', href: '#' },
        ],
    };

    return (
        <footer className="bg-gray-900 text-gray-300 pt-20 pb-10">
            <div className="max-w-7xl mx-auto px-4 md:px-8 lg:px-16">
                {/* Main Footer Content */}
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-12 mb-12">
                    {/* Brand Section */}
                    <motion.div
                        initial={{ opacity: 0, y: 20 }}
                        whileInView={{ opacity: 1, y: 0 }}
                        transition={{ duration: 0.6 }}
                    >
                        <Link href="/" className="text-2xl font-bold mb-4 inline-block">
                            HireMind<span className="text-orange-500">AI</span>
                        </Link>
                        <p className="text-gray-400">Intelligent interview platform powered by AI</p>
                    </motion.div>

                    {/* Link Sections */}
                    {Object.entries(footerLinks).map(([title, links], index) => (
                        <motion.div
                            key={title}
                            initial={{ opacity: 0, y: 20 }}
                            whileInView={{ opacity: 1, y: 0 }}
                            transition={{ duration: 0.6, delay: index * 0.1 }}
                        >
                            <h3 className="font-semibold text-white mb-4">{title}</h3>
                            <ul className="space-y-2">
                                {links.map((link) => (
                                    <li key={link.label}>
                                        <Link
                                            href={link.href}
                                            className="text-gray-400 hover:text-white transition-colors flex items-center gap-1 group"
                                        >
                                            {link.label}
                                            <HiArrowRight className="opacity-0 group-hover:opacity-100 transition-opacity" />
                                        </Link>
                                    </li>
                                ))}
                            </ul>
                        </motion.div>
                    ))}
                </div>

                {/* Bottom Bar */}
                <div className="border-t border-gray-800 pt-8 flex flex-col md:flex-row justify-between items-center gap-4">
                    <p className="text-gray-400">© {currentYear} HireMind AI. All rights reserved.</p>
                    <div className="flex gap-6">
                        {[
                            { label: 'Twitter', href: '#' },
                            { label: 'LinkedIn', href: '#' },
                            { label: 'GitHub', href: '#' },
                        ].map((social) => (
                            <Link
                                key={social.label}
                                href={social.href}
                                className="text-gray-400 hover:text-white transition-colors"
                            >
                                {social.label}
                            </Link>
                        ))}
                    </div>
                </div>
            </div>
        </footer>
    );
}
