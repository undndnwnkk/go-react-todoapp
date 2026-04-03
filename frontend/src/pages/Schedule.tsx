import React from 'react';
import { motion } from 'motion/react';
import { Clock, Coffee, Dumbbell, Briefcase, Utensils, Moon, CheckCircle2, MoreHorizontal, Bell } from 'lucide-react';
import { cn } from '../lib/utils';
import { useApp } from '../context/AppContext';
import { 
  DndContext, 
  closestCenter,
  KeyboardSensor,
  PointerSensor,
  useSensor,
  useSensors,
  DragEndEvent
} from '@dnd-kit/core';
import {
  arrayMove,
  SortableContext,
  sortableKeyboardCoordinates,
  verticalListSortingStrategy,
  useSortable
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import { RoutineItem } from '../types';

interface SortableRoutineItemProps {
  item: RoutineItem;
}

const SortableRoutineItem: React.FC<SortableRoutineItemProps> = ({ item }) => {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging
  } = useSortable({ id: item.id });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    zIndex: isDragging ? 50 : 'auto',
  };

  return (
    <div
      ref={setNodeRef}
      style={style}
      {...attributes}
      {...listeners}
      className={cn(
        "relative pl-8 mb-8 group cursor-grab active:cursor-grabbing",
        isDragging && "opacity-50"
      )}
    >
      <div className={cn(
        "absolute left-[-4px] top-1 w-2 h-2 rounded-full border-2 border-white ring-4 ring-white transition-all",
        item.isCompleted ? "bg-brand-600" : "bg-gray-300",
        isDragging && "scale-150 ring-brand-100"
      )} />
      <p className="text-xs font-bold text-gray-400 mb-2">{item.startTime} — {item.endTime}</p>
      <div className={cn(
        "p-4 rounded-xl border transition-all",
        item.isCompleted 
          ? "bg-gray-50 border-gray-100 opacity-60" 
          : "bg-white border-gray-200 shadow-sm hover:border-brand-200",
        isDragging && "shadow-2xl border-brand-500"
      )}>
        <div className="flex items-start justify-between">
          <div className="flex items-center gap-3">
            <div className="p-2 bg-brand-50 rounded-lg text-brand-600">
              {item.title.includes('Coffee') && <Coffee className="w-4 h-4" />}
              {item.title.includes('Workout') && <Dumbbell className="w-4 h-4" />}
              {item.title.includes('Interface') && <Briefcase className="w-4 h-4" />}
              {item.title.includes('Lunch') && <Utensils className="w-4 h-4" />}
              {item.title.includes('Email') && <Bell className="w-4 h-4" />}
              {item.title.includes('Walk') && <Moon className="w-4 h-4" />}
            </div>
            <div>
              <h4 className={cn("font-bold text-gray-900", item.isCompleted && "line-through")}>
                {item.title}
              </h4>
              <p className="text-xs text-gray-500 mt-1">
                {item.description || 'Routine activity'}
              </p>
            </div>
          </div>
          {item.isCompleted && <CheckCircle2 className="w-5 h-5 text-emerald-500" />}
        </div>
      </div>
    </div>
  );
};

const hours = Array.from({ length: 10 }, (_, i) => i + 6); // 6 AM to 3 PM

export const Schedule = () => {
  const { routine, updateRoutine } = useApp();
  
  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;

    if (over && active.id !== over.id) {
      const oldIndex = routine.findIndex((i) => i.id === active.id);
      const newIndex = routine.findIndex((i) => i.id === over.id);
      
      const newRoutine = arrayMove(routine, oldIndex, newIndex);
      updateRoutine(newRoutine);
    }
  };

  return (
    <div className="p-8 grid grid-cols-1 lg:grid-cols-3 gap-10">
      <div className="lg:col-span-2 space-y-8">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-4xl font-bold text-gray-900">The Flow</h1>
            <p className="text-lg text-gray-500 mt-1">
              {new Date().toLocaleDateString('en-US', { weekday: 'long', month: 'short', day: 'numeric' })} — {routine.length} Habits Active
            </p>
          </div>
          <div className="flex items-center gap-2">
            <button className="px-4 py-2 bg-gray-100 text-gray-900 rounded-xl font-semibold text-sm hover:bg-gray-200 transition-all">Today</button>
            <button className="p-2 bg-gray-100 text-gray-500 rounded-xl hover:bg-gray-200 transition-all">
              <MoreHorizontal className="w-5 h-5" />
            </button>
          </div>
        </div>

        <div className="bg-white/50 rounded-3xl border border-gray-100 p-8 relative overflow-hidden min-h-[600px]">
          <div className="absolute left-8 top-8 bottom-8 w-px bg-gray-100" />
          
          <DndContext 
            sensors={sensors}
            collisionDetection={closestCenter}
            onDragEnd={handleDragEnd}
          >
            <SortableContext 
              items={routine.map(i => i.id)}
              strategy={verticalListSortingStrategy}
            >
              <div className="relative z-10">
                {routine.map((item) => (
                  <SortableRoutineItem key={item.id} item={item} />
                ))}
              </div>
            </SortableContext>
          </DndContext>

          <div className="pt-8 border-t border-gray-100 mt-8">
            <div className="flex items-center justify-between mb-4">
              <div className="flex items-center gap-2">
                <div className="w-8 h-8 rounded-full border-4 border-emerald-500 border-r-transparent rotate-45" />
                <span className="text-sm font-bold text-gray-900">
                  {Math.round((routine.filter(r => r.isCompleted).length / routine.length) * 100)}%
                </span>
              </div>
              <div className="text-right">
                <p className="text-xs font-bold text-gray-900">Flow Consistency</p>
                <p className="text-[10px] text-gray-500">Maintain your daily rhythm to stay productive.</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="space-y-8">
        <div className="bg-white p-8 rounded-3xl border border-gray-100 shadow-sm space-y-8">
          <h2 className="text-xl font-bold text-gray-900">Daily Rhythm</h2>
          <div className="flex items-center justify-center py-6">
            <div className="relative w-48 h-48">
              <svg className="w-full h-full transform -rotate-90">
                <circle cx="96" cy="96" r="88" stroke="currentColor" strokeWidth="12" fill="transparent" className="text-gray-100" />
                <circle 
                  cx="96" 
                  cy="96" 
                  r="88" 
                  stroke="currentColor" 
                  strokeWidth="12" 
                  fill="transparent" 
                  strokeDasharray="552.92" 
                  strokeDashoffset={552.92 * (1 - routine.filter(r => r.isCompleted).length / routine.length)} 
                  className="text-emerald-500 transition-all duration-1000" 
                />
              </svg>
              <div className="absolute inset-0 flex flex-col items-center justify-center">
                <span className="text-4xl font-bold text-gray-900">
                  {Math.round((routine.filter(r => r.isCompleted).length / routine.length) * 100)}%
                </span>
                <span className="text-[10px] font-bold text-gray-400 uppercase tracking-wider">Completed</span>
              </div>
            </div>
          </div>
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <span className="text-sm font-medium text-gray-500">Focus Hours</span>
              <span className="text-sm font-bold text-brand-600">2.5 / 4h</span>
            </div>
            <div className="flex items-center justify-between">
              <span className="text-sm font-medium text-gray-500">Movement</span>
              <span className="text-sm font-bold text-emerald-600">Active</span>
            </div>
          </div>
        </div>

        <div className="bg-white p-8 rounded-3xl border border-gray-100 shadow-sm space-y-6">
          <div className="flex items-center justify-between">
            <h2 className="text-xl font-bold text-gray-900">Upcoming</h2>
            <button className="text-[10px] font-bold text-brand-600 uppercase tracking-wider hover:underline">View All</button>
          </div>
          <div className="space-y-4">
            {routine.filter(r => !r.isCompleted).slice(0, 2).map(item => (
              <div key={item.id} className="flex items-center gap-4 p-4 bg-gray-50 rounded-2xl border border-gray-100 group hover:border-brand-200 transition-all cursor-pointer">
                <div className="p-3 bg-brand-100 text-brand-600 rounded-xl group-hover:bg-brand-600 group-hover:text-white transition-all">
                  <Bell className="w-5 h-5" />
                </div>
                <div>
                  <h4 className="font-bold text-gray-900">{item.title}</h4>
                  <p className="text-[10px] font-semibold text-gray-400 uppercase tracking-wider mt-0.5">{item.startTime} • {item.type}</p>
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="relative rounded-3xl overflow-hidden group cursor-pointer">
          <img 
            src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&w=800&q=80" 
            alt="Quote" 
            className="w-full h-48 object-cover transition-transform duration-700 group-hover:scale-110"
            referrerPolicy="no-referrer"
          />
          <div className="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent p-6 flex flex-col justify-end">
            <p className="text-white font-display italic text-sm leading-relaxed">
              "Your future is found in your daily routine."
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};
