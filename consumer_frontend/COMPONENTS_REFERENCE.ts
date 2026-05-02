// Quick reference: Component exports and usage

// ============================================
// PAGES (Auto-routed by Next.js)
// ============================================

// src/app/page.tsx
// Route: /
// Exports: Home component with all sections

// src/app/about/page.tsx
// Route: /about
// Exports: About page with story, values, team

// src/app/what-we-do/page.tsx
// Route: /what-we-do
// Exports: What We Do page with features

// src/app/pricing/page.tsx
// Route: /pricing
// Exports: Pricing page with 3 tiers

// src/app/contact/page.tsx
// Route: /contact
// Exports: Contact page with form

// src/app/get-started/page.tsx
// Route: /get-started
// Exports: Get Started page with onboarding

// src/app/login/page.tsx
// Route: /login
// Exports: Login page with auth form

// ============================================
// COMMON COMPONENTS
// ============================================

import { Button } from '@/components/common/Button';
// Usage:
// <Button href="/page" variant="primary" size="lg">Click Me</Button>
// Variants: primary | secondary | outline
// Sizes: sm | md | lg

import { Navbar } from '@/components/common/Navbar';
// Usage: <Navbar />
// Auto-includes links from constants/navigation.ts

import { Footer } from '@/components/common/Footer';
// Usage: <Footer />
// Full featured footer with links

import { Section } from '@/components/common/Section';
// Usage: <Section className="bg-white"><Content /></Section>
// Animated section wrapper with scroll triggers

import { FeatureCard } from '@/components/common/FeatureCard';
// Usage:
// <FeatureCard 
//   icon="🚀"
//   title="Feature Title"
//   description="Feature description"
// />

// ============================================
// SECTION COMPONENTS
// ============================================

import { HeroSection } from '@/components/sections/HeroSection';
// Usage: <HeroSection />
// Large hero banner with CTA

import { FeaturesSection } from '@/components/sections/FeaturesSection';
// Usage: <FeaturesSection />
// 6-feature grid from content.ts

import { InterviewTypesSection } from '@/components/sections/InterviewTypesSection';
// Usage: <InterviewTypesSection />
// 3 interview type cards

import { TestimonialsSection } from '@/components/sections/TestimonialsSection';
// Usage: <TestimonialsSection />
// 3 testimonial cards

import { CTASection } from '@/components/sections/CTASection';
// Usage: <CTASection />
// Call-to-action section

// ============================================
// UTILITIES
// ============================================

import { cn } from '@/utils/cn';
// Combine classnames safely
// Usage: cn('text-blue-600', condition && 'font-bold')

import {
    fadeIn,
    fadeInUp,
    fadeInDown,
    slideInLeft,
    slideInRight,
    scaleIn,
    staggerContainer,
    staggerItem,
} from '@/utils/animations';
// Framer Motion animation variants
// Usage:
// <motion.div
//   initial="hidden"
//   whileInView="visible"
//   variants={fadeInUp}
// >
//   Content
// </motion.div>

// ============================================
// CONSTANTS
// ============================================

import { NAV_LINKS, CTA_BUTTONS } from '@/constants/navigation';
// NAV_LINKS: Array of navigation items
// CTA_BUTTONS: Primary and secondary CTA buttons

import {
    BRAND_NAME,
    BRAND_TAGLINE,
    FEATURES,
    INTERVIEW_TYPES,
    TESTIMONIALS,
} from '@/constants/content';
// All page content in easy-to-edit format

// ============================================
// EXAMPLE USAGE IN A COMPONENT
// ============================================

/*
'use client';

import { motion } from 'framer-motion';
import { Button } from '@/components/common/Button';
import { Section } from '@/components/common/Section';
import { fadeInUp } from '@/utils/animations';

export function MyComponent() {
  return (
    <Section className="bg-blue-50">
      <motion.div
        initial="hidden"
        whileInView="visible"
        variants={fadeInUp}
      >
        <h1>Welcome</h1>
        <Button href="/get-started" variant="primary">
          Get Started
        </Button>
      </motion.div>
    </Section>
  );
}
*/

// ============================================
// TAILWIND CLASSES REFERENCE
// ============================================

/*
Text Sizes:
- text-sm (12px)
- text-base (16px)
- text-lg (18px)
- text-xl (20px)
- text-2xl (24px)
- text-3xl (30px)
- text-4xl (36px)
- text-5xl (48px)
- text-6xl (60px)

Font Weight:
- font-normal (400)
- font-semibold (600)
- font-bold (700)

Colors:
- text-gray-{50-900}
- text-blue-{50-900}
- text-orange-{50-900}
- bg-gradient-to-r from-blue-600 to-orange-500

Spacing:
- p-{1-12}: padding
- m-{1-12}: margin
- gap-{1-12}: gap between flex items

Responsive:
- sm: (640px)
- md: (768px)
- lg: (1024px)
- xl: (1280px)

Usage:
- md:text-4xl (4xl on medium screens and up)
- hidden md:flex (hidden by default, flex on md+)
*/

// ============================================
// FOLDER STRUCTURE FOR NEW COMPONENTS
// ============================================

/*
To add a new component:

1. Common component (reusable):
   src/components/common/MyComponent.tsx

2. Page-specific component:
   src/components/sections/MySection.tsx

3. Import pattern:
   import { MyComponent } from '@/components/common/MyComponent';

4. Make sure to add 'use client' at top if it uses hooks/events

5. Export as default or named export
*/

// ============================================
// ENVIRONMENT VARIABLES
// ============================================

/*
Create .env.local in project root:

NEXT_PUBLIC_API_URL=https://api.yourdomain.com
NEXT_PUBLIC_STRIPE_KEY=pk_test_...

Access in components:
const apiUrl = process.env.NEXT_PUBLIC_API_URL;
*/

// ============================================
// STYLING BEST PRACTICES
// ============================================

/*
1. Use Tailwind classes for styling
   ✅ <div className="text-blue-600 hover:text-blue-700">
   ❌ <div style={{ color: 'blue' }}>

2. Use responsive prefixes
   ✅ <h1 className="text-4xl md:text-5xl lg:text-6xl">
   ❌ <h1 className="text-6xl">

3. Group related classes
   ✅ className="flex items-center gap-4 px-6 py-3 bg-blue-600 text-white rounded-lg"

4. Use Framer Motion for animations
   ✅ <motion.div whileHover={{ scale: 1.05 }} />
   ❌ <div style={{ transition: 'all 0.3s' }} />
*/

export { };
