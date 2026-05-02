'use client';

import { motion } from 'framer-motion';
import { fadeInUp, staggerContainer, staggerItem } from '@/utils/animations';
import { HiMail, HiPhone, HiLocationMarker } from 'react-icons/hi';

export default function Contact() {
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
                        Get in <span className="text-blue-600">Touch</span>
                    </h1>
                    <p className="text-xl text-gray-600 max-w-2xl mx-auto">
                        Have questions? Our team is here to help you succeed.
                    </p>
                </div>
            </motion.section>

            {/* Contact Content */}
            <section className="py-20 px-4">
                <div className="max-w-6xl mx-auto">
                    <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
                        {/* Contact Form */}
                        <motion.div
                            initial="hidden"
                            whileInView="visible"
                            viewport={{ once: true }}
                            variants={staggerContainer}
                        >
                            <h2 className="text-3xl font-bold text-gray-900 mb-8">Send us a Message</h2>

                            <form className="space-y-6">
                                {/* Name */}
                                <motion.div variants={staggerItem}>
                                    <label className="block text-sm font-semibold text-gray-900 mb-2">
                                        Full Name
                                    </label>
                                    <input
                                        type="text"
                                        placeholder="John Doe"
                                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                                    />
                                </motion.div>

                                {/* Email */}
                                <motion.div variants={staggerItem}>
                                    <label className="block text-sm font-semibold text-gray-900 mb-2">
                                        Email Address
                                    </label>
                                    <input
                                        type="email"
                                        placeholder="john@example.com"
                                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                                    />
                                </motion.div>

                                {/* Subject */}
                                <motion.div variants={staggerItem}>
                                    <label className="block text-sm font-semibold text-gray-900 mb-2">
                                        Subject
                                    </label>
                                    <input
                                        type="text"
                                        placeholder="How can we help?"
                                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                                    />
                                </motion.div>

                                {/* Message */}
                                <motion.div variants={staggerItem}>
                                    <label className="block text-sm font-semibold text-gray-900 mb-2">
                                        Message
                                    </label>
                                    <textarea
                                        rows={6}
                                        placeholder="Tell us more about your inquiry..."
                                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-600"
                                    ></textarea>
                                </motion.div>

                                {/* Submit Button */}
                                <motion.button
                                    variants={staggerItem}
                                    type="submit"
                                    className="w-full bg-gradient-to-r from-blue-600 to-orange-500 text-white font-bold py-3 rounded-lg hover:shadow-lg transition-all"
                                    whileHover={{ scale: 1.02 }}
                                    whileTap={{ scale: 0.98 }}
                                >
                                    Send Message
                                </motion.button>
                            </form>
                        </motion.div>

                        {/* Contact Info */}
                        <motion.div
                            initial="hidden"
                            whileInView="visible"
                            viewport={{ once: true }}
                            variants={staggerContainer}
                        >
                            <h2 className="text-3xl font-bold text-gray-900 mb-8">Contact Information</h2>

                            <div className="space-y-8">
                                {/* Email */}
                                <motion.div
                                    variants={staggerItem}
                                    whileHover={{ x: 10 }}
                                    className="flex gap-4 items-start p-6 bg-gray-50 rounded-xl border border-gray-200 hover:border-blue-400 transition-all"
                                >
                                    <div className="text-3xl text-blue-600">
                                        <HiMail />
                                    </div>
                                    <div>
                                        <h3 className="text-lg font-bold text-gray-900 mb-1">Email</h3>
                                        <p className="text-gray-600">support@hiremindai.com</p>
                                        <p className="text-sm text-gray-500 mt-1">We'll respond within 24 hours</p>
                                    </div>
                                </motion.div>

                                {/* Phone */}
                                <motion.div
                                    variants={staggerItem}
                                    whileHover={{ x: 10 }}
                                    className="flex gap-4 items-start p-6 bg-gray-50 rounded-xl border border-gray-200 hover:border-blue-400 transition-all"
                                >
                                    <div className="text-3xl text-blue-600">
                                        <HiPhone />
                                    </div>
                                    <div>
                                        <h3 className="text-lg font-bold text-gray-900 mb-1">Phone</h3>
                                        <p className="text-gray-600">+1 (555) 123-4567</p>
                                        <p className="text-sm text-gray-500 mt-1">Monday - Friday, 9AM - 5PM EST</p>
                                    </div>
                                </motion.div>

                                {/* Location */}
                                <motion.div
                                    variants={staggerItem}
                                    whileHover={{ x: 10 }}
                                    className="flex gap-4 items-start p-6 bg-gray-50 rounded-xl border border-gray-200 hover:border-blue-400 transition-all"
                                >
                                    <div className="text-3xl text-blue-600">
                                        <HiLocationMarker />
                                    </div>
                                    <div>
                                        <h3 className="text-lg font-bold text-gray-900 mb-1">Location</h3>
                                        <p className="text-gray-600">San Francisco, CA</p>
                                        <p className="text-gray-600">New York, NY</p>
                                    </div>
                                </motion.div>
                            </div>

                            {/* Social Links */}
                            <motion.div
                                variants={staggerItem}
                                className="mt-12 pt-8 border-t border-gray-200"
                            >
                                <h3 className="text-lg font-bold text-gray-900 mb-4">Follow Us</h3>
                                <div className="flex gap-4">
                                    {['Twitter', 'LinkedIn', 'GitHub'].map((social) => (
                                        <a
                                            key={social}
                                            href="#"
                                            className="w-10 h-10 rounded-full bg-blue-100 text-blue-600 hover:bg-blue-600 hover:text-white transition-all flex items-center justify-center"
                                        >
                                            {social[0]}
                                        </a>
                                    ))}
                                </div>
                            </motion.div>
                        </motion.div>
                    </div>
                </div>
            </section>
        </main>
    );
}
