# CarPeek Project Context

## Project Overview

CarPeek is a Next.js web application called "Car Peek | Daily Car Challenge" - a daily car identification game where players guess vehicles from their tail lights. The app presents users with close-up images of car tail lights and challenges them to identify the make, model, and year of the vehicle.

### Technology Stack
- **Framework**: Next.js 16.0.10 (App Router)
- **Language**: TypeScript
- **Runtime**: Node.js
- **Styling**: CSS Modules with custom design system
- **Icons**: Lucide React
- **Deployment**: Docker container with Nginx (static export)

### Architecture
- **Frontend**: Next.js application in the `/frontend` directory
- **UI Components**: Custom component library in `/app/components/ui`
- **Game Logic**: Centralized in `/app/components/GameInterface.tsx`
- **Design System**: Custom dark-themed "Night Drive" design in `/app/globals.css`

## Key Features

1. **Daily Car Challenge**: Users identify cars from tail light images
2. **Multi-field Input**: Separate fields for Make, Model, and Year/Generation
3. **Guess Tracking**: Shows previous attempts with correctness indicators
4. **Responsive Design**: Mobile-first approach with 480px max-width container
5. **Dark Theme**: Custom "Night Drive" theme with cherry/carbon/red color scheme

## File Structure

```
frontend/
├── app/                    # Next.js App Router pages
│   ├── components/         # Reusable UI components
│   │   └── ui/            # Base UI elements (Button, Input, Header)
│   ├── globals.css        # Custom design system and theme
│   ├── layout.tsx         # Root layout with metadata
│   └── page.tsx           # Main game interface page
├── public/                # Static assets
├── Dockerfile             # Multi-stage build with Nginx
├── next.config.ts         # Next.js configuration (static export)
├── package.json           # Dependencies and scripts
└── tsconfig.json          # TypeScript configuration
```

## Development Commands

- `npm run dev` - Start development server on http://localhost:3000
- `npm run build` - Build the application for production
- `npm run start` - Start production server
- `npm run lint` - Run ESLint for code quality checks

## Build & Deployment

The application is configured for static export (`output: 'export'`) in `next.config.ts`, making it suitable for CDN hosting. The Dockerfile implements a multi-stage build:
1. Builder stage: Installs dependencies and builds static files
2. Production stage: Serves static files with Nginx

## Design System

The "Night Drive" theme uses a dark color palette with cherry/carbon/red accents:
- Primary: `#F51B40` (vibrant red/pink)
- Background: `#13090B` (deep dark brown/black)
- Surface: `#1E1114` (slightly lighter panels)
- Text: White with muted variants

## Development Conventions

- Component-based architecture with reusable UI elements
- Client-side state management with React hooks
- CSS variables for consistent theming
- Responsive design with mobile-first approach
- Semantic HTML and accessibility considerations
- TypeScript for type safety

## Environment Configuration

The project includes `.env.local` for local environment variables, though specific variables weren't visible in the scan. The Dockerfile suggests the app is designed to work in containerized environments.

## Special Notes

- Images are unoptimized due to static export requirements
- Remote patterns are configured for Unsplash images
- Custom Google Font (Outfit) is integrated
- Animation effects are implemented with CSS keyframes