import React, { createContext, useContext, useState, useEffect } from 'react';
import { Task, RoutineItem, User } from '../types';

interface AppContextType {
  user: User | null;
  tasks: Task[];
  routine: RoutineItem[];
  login: (email: string, name: string) => void;
  logout: () => void;
  addTask: (task: Omit<Task, 'id'>) => void;
  updateTask: (id: string, updates: Partial<Task>) => void;
  deleteTask: (id: string) => void;
  updateRoutine: (routine: RoutineItem[]) => void;
  updateRoutineItem: (id: string, updates: Partial<RoutineItem>) => void;
}

const AppContext = createContext<AppContextType | undefined>(undefined);

const DEFAULT_TASKS: Task[] = [
  {
    id: '1',
    title: 'Q4 Financial Audit Review',
    description: 'Complete the comprehensive review of the Q4 financial statements and milestones.',
    dueDate: '2023-10-24',
    priority: 'Critical',
    status: 'Past Due',
    category: 'Finance'
  },
  {
    id: '2',
    title: 'Brand Identity System',
    description: 'Develop the visual language for the Orchestra rebrand including color...',
    dueDate: '2023-11-12',
    priority: 'High',
    status: 'Active',
    category: 'Design'
  }
];

const DEFAULT_ROUTINE: RoutineItem[] = [
  { id: '1', title: 'Morning Coffee', startTime: '06:15', endTime: '06:45', isCompleted: true, type: 'routine', description: 'Quiet planning and ritual establishment.' },
  { id: '2', title: 'Morning Workout', startTime: '07:00', endTime: '08:00', isCompleted: true, type: 'routine', description: 'High-intensity interval training.' },
  { id: '3', title: 'Interface Orchestration', startTime: '09:00', endTime: '11:00', isCompleted: false, type: 'deep-work', description: 'Focus Mode: Engaged. No Interruptions.' },
];

export const AppProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(() => {
    const saved = localStorage.getItem('orchestra_user');
    return saved ? JSON.parse(saved) : null;
  });

  const [tasks, setTasks] = useState<Task[]>(() => {
    const saved = localStorage.getItem('orchestra_tasks');
    return saved ? JSON.parse(saved) : DEFAULT_TASKS;
  });

  const [routine, setRoutine] = useState<RoutineItem[]>(() => {
    const saved = localStorage.getItem('orchestra_routine');
    return saved ? JSON.parse(saved) : DEFAULT_ROUTINE;
  });

  useEffect(() => {
    localStorage.setItem('orchestra_user', JSON.stringify(user));
  }, [user]);

  useEffect(() => {
    localStorage.setItem('orchestra_tasks', JSON.stringify(tasks));
  }, [tasks]);

  useEffect(() => {
    localStorage.setItem('orchestra_routine', JSON.stringify(routine));
  }, [routine]);

  const login = (email: string, name: string) => {
    setUser({ id: '1', name, email, avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${name}` });
  };

  const logout = () => {
    setUser(null);
  };

  const addTask = (task: Omit<Task, 'id'>) => {
    const newTask = { ...task, id: Math.random().toString(36).substr(2, 9) };
    setTasks(prev => [newTask, ...prev]);
  };

  const updateTask = (id: string, updates: Partial<Task>) => {
    setTasks(prev => prev.map(t => t.id === id ? { ...t, ...updates } : t));
  };

  const deleteTask = (id: string) => {
    setTasks(prev => prev.filter(t => t.id !== id));
  };

  const updateRoutine = (newRoutine: RoutineItem[]) => {
    setRoutine(newRoutine);
  };

  const updateRoutineItem = (id: string, updates: Partial<RoutineItem>) => {
    setRoutine(prev => prev.map(item => item.id === id ? { ...item, ...updates } : item));
  };

  return (
    <AppContext.Provider value={{ 
      user, tasks, routine, login, logout, addTask, updateTask, deleteTask, updateRoutine, updateRoutineItem 
    }}>
      {children}
    </AppContext.Provider>
  );
};

export const useApp = () => {
  const context = useContext(AppContext);
  if (!context) throw new Error('useApp must be used within an AppProvider');
  return context;
};
