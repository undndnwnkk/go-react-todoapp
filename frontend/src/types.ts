export type Priority = 'Low' | 'Medium' | 'High' | 'Critical';
export type TaskStatus = 'Active' | 'Completed' | 'Past Due' | 'Overdue';

export interface Task {
  id: string;
  title: string;
  description: string;
  dueDate: string;
  priority: Priority;
  status: TaskStatus;
  category: string;
}

export interface RoutineItem {
  id: string;
  title: string;
  startTime: string;
  endTime: string;
  isCompleted: boolean;
  type: 'deep-work' | 'routine' | 'break';
  description?: string;
}

export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string;
}
