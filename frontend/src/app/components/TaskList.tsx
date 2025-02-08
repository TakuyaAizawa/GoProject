import { Task } from '../types';

interface TaskListProps {
  tasks: Task[];
  editingTask: Task | null;
  onStartEdit: (task: Task) => void;
  onUpdate: (e: React.FormEvent) => Promise<void>;
  onCancelEdit: () => void;
  onDelete: (id: number) => Promise<void>;
  onEditingTaskChange: (task: Task) => void;
}

export const TaskList = ({
  tasks,
  editingTask,
  onStartEdit,
  onUpdate,
  onCancelEdit,
  onDelete,
  onEditingTaskChange,
}: TaskListProps) => {
  return (
    <div className="space-y-4">
      {tasks.map(task => (
        <div 
          key={task.id} 
          className="p-4 bg-white rounded-lg shadow"
        >
          {editingTask?.id === task.id ? (
            <form onSubmit={onUpdate} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700">タイトル</label>
                <input
                  type="text"
                  value={editingTask.title}
                  onChange={(e) => onEditingTaskChange({...editingTask, title: e.target.value})}
                  className="mt-1 w-full p-2 border rounded-md"
                  required
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">説明</label>
                <input
                  type="text"
                  value={editingTask.description}
                  onChange={(e) => onEditingTaskChange({...editingTask, description: e.target.value})}
                  className="mt-1 w-full p-2 border rounded-md"
                  required
                />
              </div>
              <div className="flex gap-2">
                <button 
                  type="submit"
                  className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
                >
                  保存
                </button>
                <button 
                  type="button"
                  onClick={onCancelEdit}
                  className="px-4 py-2 border rounded-md hover:bg-gray-100"
                >
                  キャンセル
                </button>
              </div>
            </form>
          ) : (
            <div>
              <div className="flex justify-between items-start">
                <div>
                  <h2 className="font-bold text-lg">{task.title}</h2>
                  <p className="text-gray-600 mt-2">{task.description}</p>
                </div>
                <div className="flex gap-2">
                  <button
                    onClick={() => onStartEdit(task)}
                    className="px-3 py-1 text-blue-600 hover:bg-blue-50 rounded-md"
                  >
                    編集
                  </button>
                  <button
                    onClick={() => onDelete(task.id)}
                    className="px-3 py-1 text-red-600 hover:bg-red-50 rounded-md"
                  >
                    削除
                  </button>
                </div>
              </div>
            </div>
          )}
        </div>
      ))}
    </div>
  );
};
