import React, { useState } from 'react';
import { motion } from 'motion/react';
import { Plus, MoreVertical, Clock, Zap, Trash2, Edit2, CheckCircle2 } from 'lucide-react';
import { cn } from '../lib/utils';
import { useApp } from '../context/AppContext';
import { CreateTaskModal } from '../components/CreateTaskModal';
import { TaskStatus } from '../types';

const filters: (TaskStatus | 'All')[] = ['All', 'Active', 'Completed', 'Past Due'];

export const Tasks = () => {
  const { tasks, updateTask, deleteTask } = useApp();
  const [activeFilter, setActiveFilter] = useState<TaskStatus | 'All'>('All');
  const [isModalOpen, setIsModalOpen] = useState(false);

  const filteredTasks = tasks.filter(task => 
    activeFilter === 'All' ? true : task.status === activeFilter
  );

  return (
    <div className="p-8 space-y-8">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-4xl font-bold text-gray-900">Workspace Tasks</h1>
          <p className="text-lg text-gray-500 mt-1">Organize your deep work sessions and milestones.</p>
        </div>
        
        <div className="flex items-center gap-2 p-1 bg-gray-100 rounded-xl">
          {filters.map((filter) => (
            <button
              key={filter}
              onClick={() => setActiveFilter(filter)}
              className={cn(
                "px-4 py-2 rounded-lg text-sm font-semibold transition-all",
                activeFilter === filter ? "bg-white text-gray-900 shadow-sm" : "text-gray-500 hover:text-gray-900"
              )}
            >
              {filter}
            </button>
          ))}
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {filteredTasks.map((task, idx) => (
          <motion.div
            key={task.id}
            initial={{ opacity: 0, scale: 0.95 }}
            animate={{ opacity: 1, scale: 1 }}
            transition={{ delay: idx * 0.05 }}
            className={cn(
              "p-6 rounded-2xl border transition-all group relative",
              task.status === 'Past Due' ? "border-rose-200 bg-rose-50/30" : "bg-white border-gray-100 shadow-sm hover:border-brand-200"
            )}
          >
            <div className="flex items-start justify-between mb-4">
              <span className={cn(
                "px-2 py-1 rounded-md text-[10px] font-bold uppercase tracking-wider",
                task.status === 'Past Due' ? "bg-rose-100 text-rose-600" : 
                task.status === 'Completed' ? "bg-emerald-100 text-emerald-600" : "bg-brand-50 text-brand-600"
              )}>
                {task.status}
              </span>
              <div className="flex items-center gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                <button 
                  onClick={() => updateTask(task.id, { status: task.status === 'Completed' ? 'Active' : 'Completed' })}
                  className="p-1.5 text-gray-400 hover:text-brand-600 hover:bg-brand-50 rounded-md"
                >
                  <CheckCircle2 className="w-4 h-4" />
                </button>
                <button 
                  onClick={() => deleteTask(task.id)}
                  className="p-1.5 text-gray-400 hover:text-rose-600 hover:bg-rose-50 rounded-md"
                >
                  <Trash2 className="w-4 h-4" />
                </button>
              </div>
            </div>

            <h3 className={cn(
              "text-xl font-bold text-gray-900 mb-2",
              task.status === 'Completed' && "line-through opacity-60"
            )}>
              {task.title}
            </h3>
            <p className="text-gray-500 text-sm mb-6 line-clamp-2">{task.description}</p>

            <div className="flex items-center justify-between mt-auto">
              <div className="flex items-center gap-4 text-xs font-semibold text-gray-400">
                <div className="flex items-center gap-1.5">
                  <Clock className="w-3.5 h-3.5" />
                  <span className={task.status === 'Past Due' ? "text-rose-600" : ""}>{task.dueDate}</span>
                </div>
                <div className="flex items-center gap-1.5">
                  <Zap className={cn("w-3.5 h-3.5", task.priority === 'High' || task.priority === 'Critical' ? "text-brand-500" : "text-gray-400")} />
                  <span>{task.priority}</span>
                </div>
              </div>
              {task.status === 'Completed' && (
                <div className="w-6 h-6 rounded-full bg-emerald-500 flex items-center justify-center">
                  <CheckCircle2 className="w-4 h-4 text-white" />
                </div>
              )}
            </div>
          </motion.div>
        ))}

        <motion.button
          whileHover={{ scale: 1.02 }}
          whileTap={{ scale: 0.98 }}
          onClick={() => setIsModalOpen(true)}
          className="p-6 rounded-2xl border-2 border-dashed border-gray-200 flex flex-col items-center justify-center gap-4 text-gray-400 hover:border-brand-300 hover:text-brand-500 transition-all group"
        >
          <div className="w-12 h-12 rounded-full bg-gray-50 flex items-center justify-center group-hover:bg-brand-50 transition-all">
            <Plus className="w-6 h-6" />
          </div>
          <div className="text-center">
            <p className="font-bold text-gray-900">Add New Task</p>
            <p className="text-sm">Start a new workflow</p>
          </div>
        </motion.button>
      </div>
      <CreateTaskModal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} />
    </div>
  );
};
