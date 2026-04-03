# Orchestra - Task & Routine Manager

Orchestra is a professional task management and daily routine orchestration application designed for high-performance individuals. It combines traditional task listing with time-blocked daily scheduling to help users achieve a state of "Deep Work" and maintain a balanced flow throughout the day.

## 🚀 Key Features

- **Personalized Dashboard**: Get a high-level overview of your total scope, completed focus sessions, and tasks requiring immediate attention.
- **Priority Focus**: Intelligent task prioritization that highlights what matters most right now.
- **Dynamic Task Management**: Organize tasks by category, due date, and priority levels (Low to Critical).
- **Daily Routine Orchestration**: A visual timeline for your day, allowing you to block time for Deep Work, routine habits, and mindful breaks.
- **Daily Rhythm Analytics**: Track your consistency and focus hours with real-time progress visualization.
- **Modern UI/UX**: Built with a "glassmorphism" aesthetic, featuring smooth transitions, responsive layouts, and professional typography.

## 🛠 Tech Stack

- **Frontend**: React 19 with TypeScript
- **Styling**: Tailwind CSS 4.0
- **Animations**: Motion (formerly Framer Motion)
- **Icons**: Lucide React
- **Routing**: React Router 6
- **Date Handling**: date-fns
- **Build Tool**: Vite

## 📂 Project Structure

```text
src/
├── components/       # Reusable UI components (Layout, Sidebar, Header)
├── lib/              # Utility functions (cn helper, formatting)
├── pages/            # Main application views (Dashboard, Tasks, Schedule, Login)
├── types.ts          # TypeScript interfaces and enums
├── App.tsx           # Main routing and protected route logic
└── index.css         # Global styles and Tailwind theme configuration
```

## 🚦 Getting Started

### Prerequisites
- Node.js (Latest LTS recommended)
- npm or yarn

### Installation
1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm run dev
   ```

## 📖 Usage Guide

### Authentication
The application currently uses a **Mock Authentication** system for demonstration purposes. 
- Navigate to the `/login` page.
- Enter any email and password to access the dashboard.
- User data is persisted in `localStorage`.

### Managing Tasks
- Go to the **Tasks** tab to view your current workspace.
- Tasks are color-coded based on their status:
  - **Active**: Blue
  - **Completed**: Emerald (with strikethrough)
  - **Past Due**: Rose (with warning background)

### Planning Your Flow
- Use the **Schedule** tab to visualize your daily rhythm.
- The timeline starts from 6:00 AM and shows your planned "Flow".
- **Deep Work** blocks are highlighted in solid brand colors to signify intense focus periods.

## 🎨 Design System

Orchestra uses a custom theme defined in `src/index.css`:
- **Font Sans**: Inter (for readability)
- **Font Display**: Outfit (for headings)
- **Brand Colors**: A custom blue scale (`#6b8ef7`) representing trust and clarity.

---
*Created with Google AI Studio Build.*
