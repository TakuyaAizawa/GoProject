import { Todo } from '../types';

interface TodoListProps {
  todos: Todo[];
  onStartEdit: (id: number) => void;
  onUpdate: (id: number, text: string) => Promise<void>;
  onCancelEdit: (id: number) => void;
  onDelete: (id: number) => Promise<void>;
}

export const TodoList = ({
  todos,
  onStartEdit,
  onUpdate,
  onCancelEdit,
  onDelete,
}: TodoListProps) => {
  return (
    <ul className="space-y-2">
      {todos.map((todo) => (
        <li 
          key={todo.id} 
          className="p-3 bg-white rounded-lg shadow"
        >
          {todo.isEditing ? (
            <div className="flex gap-2">
              <input
                type="text"
                defaultValue={todo.text}
                className="flex-1 p-2 border rounded-md"
                onKeyDown={(e) => {
                  if (e.key === 'Enter') {
                    onUpdate(todo.id, e.currentTarget.value)
                  } else if (e.key === 'Escape') {
                    onCancelEdit(todo.id)
                  }
                }}
                autoFocus
              />
              <button
                onClick={() => onCancelEdit(todo.id)}
                className="px-3 py-1 border rounded-md hover:bg-gray-100"
              >
                キャンセル
              </button>
            </div>
          ) : (
            <div className="flex items-center justify-between">
              <span className="flex-1">{todo.text}</span>
              <div className="flex gap-2">
                <button
                  onClick={() => onStartEdit(todo.id)}
                  className="px-3 py-1 text-blue-600 hover:bg-blue-50 rounded-md"
                >
                  編集
                </button>
                <button
                  onClick={() => onDelete(todo.id)}
                  className="px-3 py-1 text-red-600 hover:bg-red-50 rounded-md"
                >
                  削除
                </button>
              </div>
            </div>
          )}
        </li>
      ))}
    </ul>
  );
};
