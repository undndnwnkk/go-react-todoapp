import React, { useState } from 'react';
import { motion } from 'motion/react';
import { CheckCircle, AlertCircle, BarChart3, Plus, ChevronRight, Clock } from 'lucide-react';
import { cn } from '../lib/utils';
import { useApp } from '../context/AppContext';
import { CreateTaskModal } from '../components/CreateTaskModal';

export const Dashboard = () => {
  const { tasks, routine, user, updateTask } = useApp();
  const [isModalOpen, setIsModalOpen] = useState(false);

  const activeTasks = tasks.filter(t => t.status !== 'Completed');
  const completedTasks = tasks.filter(t => t.status === 'Completed');
  const overdueTasks = tasks.filter(t => t.status === 'Past Due');

  const stats = [
    { label: 'Total Scope', value: tasks.length.toString().padStart(2, '0'), icon: BarChart3, color: 'bg-brand-50 text-brand-600' },
    { label: 'Completed Focus', value: completedTasks.length.toString().padStart(2, '0'), icon: CheckCircle, color: 'bg-emerald-50 text-emerald-600' },
    { label: 'Requires Attention', value: overdueTasks.length.toString().padStart(2, '0'), icon: AlertCircle, color: 'bg-rose-50 text-rose-600' },
  ];

  return (
    <div className="p-8 space-y-10">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-4xl font-bold text-gray-900">Good morning, {user?.name || 'Julian'}.</h1>
          <p className="text-lg text-gray-500 mt-1">Your flow is looking balanced today.</p>
        </div>
        <button 
          onClick={() => setIsModalOpen(true)}
          className="flex items-center gap-2 px-6 py-3 bg-brand-600 text-white rounded-xl font-semibold hover:bg-brand-700 transition-all shadow-lg shadow-brand-200"
        >
          <Plus className="w-5 h-5" />
          Create Task
        </button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {stats.map((stat, idx) => (
          <motion.div
            key={stat.label}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: idx * 0.1 }}
            className="p-6 bg-white rounded-2xl border border-gray-100 shadow-sm flex items-center justify-between group hover:border-brand-200 transition-all"
          >
            <div>
              <p className="text-sm font-medium text-gray-500 mb-2">{stat.label}</p>
              <h3 className="text-3xl font-bold text-gray-900">{stat.value}</h3>
            </div>
            <div className={cn("p-3 rounded-xl transition-transform group-hover:scale-110", stat.color)}>
              <stat.icon className="w-6 h-6" />
            </div>
          </motion.div>
        ))}
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2 space-y-6">
          <div className="flex items-center justify-between">
            <h2 className="text-xl font-bold text-gray-900">Priority Focus</h2>
            <button className="text-sm font-semibold text-brand-600 hover:text-brand-700 flex items-center gap-1">
              View All <ChevronRight className="w-4 h-4" />
            </button>
          </div>

          <div className="space-y-4">
            {activeTasks.slice(0, 3).map((task) => (
              <motion.div
                key={task.id}
                whileHover={{ x: 4 }}
                className="p-6 bg-white rounded-2xl border border-gray-100 shadow-sm hover:border-brand-200 transition-all group"
              >
                <div className="flex items-start justify-between mb-4">
                  <div className="flex items-center gap-3">
                    <button 
                      onClick={() => updateTask(task.id, { status: 'Completed' })}
                      className="w-6 h-6 rounded-full border-2 border-gray-200 group-hover:border-brand-500 transition-all cursor-pointer" 
                    />
                    <h3 className="text-lg font-bold text-gray-900">{task.title}</h3>
                  </div>
                  <span className={cn(
                    "px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wider",
                    task.priority === 'High' || task.priority === 'Critical' ? "bg-brand-50 text-brand-600" : "bg-gray-100 text-gray-600"
                  )}>
                    {task.priority} Priority
                  </span>
                </div>
                <p className="text-gray-500 mb-6 line-clamp-2">{task.description}</p>
                <div className="flex items-center gap-6 text-sm text-gray-400">
                  <div className="flex items-center gap-2">
                    <Clock className="w-4 h-4" />
                    <span>{task.dueDate}</span>
                  </div>
                  <div className="flex items-center gap-2">
                    <div className="w-2 h-2 rounded-full bg-brand-400" />
                    <span>{task.category}</span>
                  </div>
                </div>
              </motion.div>
            ))}
            {completedTasks.slice(0, 2).map(task => (
              <motion.div
                key={task.id}
                whileHover={{ x: 4 }}
                className="p-6 bg-gray-50/50 rounded-2xl border border-dashed border-gray-200 flex items-center gap-4 opacity-60"
              >
                <div className="w-6 h-6 rounded-full bg-emerald-500 flex items-center justify-center">
                  <CheckCircle className="w-4 h-4 text-white" />
                </div>
                <h3 className="text-lg font-bold text-gray-400 line-through">{task.title}</h3>
              </motion.div>
            ))}
          </div>
        </div>

        <div className="space-y-6">
          <div className="flex items-center justify-between">
            <h2 className="text-xl font-bold text-gray-900">Daily Routine</h2>
            <span className="text-xs font-bold text-brand-600 bg-brand-50 px-2 py-1 rounded-md uppercase tracking-wider">
              {new Date().toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })}
            </span>
          </div>

          <div className="bg-white p-6 rounded-2xl border border-gray-100 shadow-sm space-y-8 relative overflow-hidden">
            <div className="absolute left-8 top-8 bottom-8 w-px bg-gray-100" />
            
            {routine.slice(0, 4).map((item) => (
              <div key={item.id} className="relative pl-8">
                <div className={cn(
                  "absolute left-[-4px] top-1 w-2 h-2 rounded-full border-2 border-white ring-4 ring-white",
                  item.isCompleted ? "bg-brand-600" : "bg-gray-300"
                )} />
                <p className="text-xs font-bold text-gray-400 mb-2">{item.startTime} — {item.endTime}</p>
                <div className={cn(
                  "p-4 rounded-xl border transition-all",
                  item.isCompleted 
                    ? "bg-gray-50 border-gray-100 opacity-60" 
                    : "bg-white border-gray-200 shadow-sm hover:border-brand-200"
                )}>
                  <h4 className={cn("font-bold text-gray-900", item.isCompleted && "line-through")}>
                    {item.title}
                  </h4>
                  <p className="text-xs text-gray-500 mt-1">
                    {item.description || 'Quiet planning and ritual establishment.'}
                  </p>
                </div>
              </div>
            ))}

            <div className="pt-4 border-t border-gray-100">
              <div className="flex items-center justify-between mb-4">
                <div className="flex items-center gap-2">
                  <div className="w-8 h-8 rounded-full border-4 border-emerald-500 border-r-transparent rotate-45" />
                  <span className="text-sm font-bold text-gray-900">
                    {Math.round((routine.filter(r => r.isCompleted).length / routine.length) * 100)}%
                  </span>
                </div>
                <div className="text-right">
                  <p className="text-xs font-bold text-gray-900">Flow Consistency</p>
                  <p className="text-[10px] text-gray-500">You've maintained focus for 4 hours today.</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <CreateTaskModal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} />
    </div>
  );
};
