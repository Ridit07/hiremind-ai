# HireMind AI - Consumer Frontend

A modern, production-grade frontend for the HireMind AI interview preparation platform built with Next.js 14+, TypeScript, Tailwind CSS, and Framer Motion.

## 🚀 Quick Start

### Prerequisites
- Node.js 18+ 
- npm or yarn

### Installation

```bash
# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build

# Start production server
npm run start
```

Visit `http://localhost:3000` to see the application.

## 📁 Project Structure

```
src/
├── app/                    # Next.js app router pages
│   ├── page.tsx           # Home page
│   ├── about/page.tsx     # About us page
│   ├── what-we-do/page.tsx # What we do page
│   ├── pricing/page.tsx    # Pricing page
│   ├── contact/page.tsx    # Contact page
│   ├── get-started/page.tsx # Get started/signup page
│   ├── login/page.tsx      # Login page
│   ├── layout.tsx          # Root layout
│   └── globals.css         # Global styles
├── components/
│   ├── common/             # Reusable components
│   │   ├── Button.tsx      # CTA Button component
│   │   ├── Navbar.tsx      # Navigation bar
│   │   ├── Footer.tsx      # Footer
│   │   ├── Section.tsx     # Section wrapper
│   │   └── FeatureCard.tsx # Feature card component
│   └── sections/           # Page sections
│       ├── HeroSection.tsx
│       ├── FeaturesSection.tsx
│       ├── InterviewTypesSection.tsx
│       ├── TestimonialsSection.tsx
│       └── CTASection.tsx
├── constants/              # App constants
│   ├── navigation.ts       # Navigation links
│   └── content.ts          # Page content
├── utils/                  # Utility functions
│   ├── cn.ts              # Classname utility
│   └── animations.ts      # Framer Motion animation variants
└── hooks/                 # Custom React hooks (for future use)
```

## 🎨 Design & Styling

- **Framework**: Tailwind CSS for utility-first styling
- **Animations**: Framer Motion for smooth, production-grade animations
- **Icons**: React Icons for scalable SVG icons
- **Color Scheme**: Blue (#1E40AF) & Orange (#F97316) gradient theme

## ✨ Features

### Pages Implemented

1. **Home** - Beautiful hero section with CTA, features showcase, interview types, testimonials
2. **About Us** - Company story, values, team showcase
3. **What We Do** - Detailed explanation of capabilities and technology
4. **Pricing** - Three-tier pricing plan comparison
5. **Contact** - Contact form and business information
6. **Get Started** - Onboarding process with FAQ
7. **Login** - Authentication page with social login options

### Key Components

- **Responsive Design** - Mobile-first, fully responsive across all devices
- **Animations** - Smooth fade-in, slide, and scale animations on scroll
- **Accessibility** - Semantic HTML, proper ARIA labels
- **Performance** - Image optimization, code splitting, lazy loading

## 🔧 Configuration

### Tailwind CSS
Customized in `tailwind.config.ts` with:
- Custom color palette
- Extended spacing
- Custom animations

### TypeScript
Strict mode enabled for type safety. Path aliases configured:
- `@/*` → `src/`

## 🚀 Production Build

```bash
npm run build
npm run start
```

The production build is optimized for performance with:
- Minified code
- Image optimization
- CSS optimization
- Code splitting

## 📊 Browser Support

- Chrome/Edge (latest)
- Firefox (latest)
- Safari (latest)
- Mobile browsers

## 🔐 Security

- No sensitive data hardcoded
- All forms are non-functional (for now)
- Safe external link handling
- CSP headers ready

## 📝 Next Steps

### To-Do List
- [ ] Integrate backend API for forms
- [ ] Implement authentication flow
- [ ] Add payment integration (Stripe/Razorpay)
- [ ] Analytics integration (Google Analytics, Mixpanel)
- [ ] Email service integration
- [ ] User dashboard
- [ ] Interview platform interface
- [ ] Admin dashboard

## 🤝 Contributing

1. Follow the existing code style
2. Maintain component modularity
3. Keep animations performance-optimized
4. Update README for new features

## 📄 License

Proprietary - HireMind AI

---

**Last Updated**: May 2, 2026

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
