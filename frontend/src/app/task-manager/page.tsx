'use client'

import Link from 'next/link'
import { TaskForm } from '../components/TaskForm'
import { TaskList } from '../components/TaskList'
import { useTasks } from '../hooks/useTasks'

export default function TaskManager() {
  const {
    tasks,
    editingTask,
    setEditingTask,
    createTask,
    updateTask,
    deleteTask,
  } = useTasks();

  const handleCreateTask = async (title: string, description: string) => {
    await createTask(title, description);
  };

  const handleUpdateTask = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!editingTask) return;
    await updateTask(editingTask);
  };

  const handleDeleteTask = async (id: number) => {
    if (!confirm('このタスクを削除してもよろしいですか？')) return;
    await deleteTask(id);
  };

  return (
    <main className="max-w-4xl mx-auto p-4">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-2xl font-bold">タスク管理</h1>
        <Link 
          href="/tasks" 
          className="px-4 py-2 border rounded-md hover:bg-gray-100 transition-colors"
        >
          TODOリストへ→
        </Link>
      </div>
      
      <TaskForm onSubmit={handleCreateTask} />

      <TaskList
        tasks={tasks}
        editingTask={editingTask}
        onStartEdit={setEditingTask}
        onUpdate={handleUpdateTask}
        onCancelEdit={() => setEditingTask(null)}
        onDelete={handleDeleteTask}
        onEditingTaskChange={setEditingTask}
      />
    </main>
  );
}
