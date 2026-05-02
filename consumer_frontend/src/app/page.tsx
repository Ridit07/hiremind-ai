import type { Metadata } from 'next';
import { HeroSection } from '@/components/sections/HeroSection';
import { FeaturesSection } from '@/components/sections/FeaturesSection';
import { InterviewTypesSection } from '@/components/sections/InterviewTypesSection';
import { TestimonialsSection } from '@/components/sections/TestimonialsSection';
import { CTASection } from '@/components/sections/CTASection';

export const metadata: Metadata = {
  title: 'HireMind AI - Master Your Interview Skills',
  description:
    'Practice with AI-powered interviews, get real-time feedback, and land your dream job. Experience interviews like never before.',
  keywords: [
    'interview preparation',
    'AI interviews',
    'job interview practice',
    'mock interviews',
    'interview coaching',
  ],
};

export default function Home() {
  return (
    <main>
      <HeroSection />
      <FeaturesSection />
      <InterviewTypesSection />
      <TestimonialsSection />
      <CTASection />
    </main>
  );
}
