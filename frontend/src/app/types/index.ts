export interface Task {
  id: number;
  title: string;
  description: string;
}

export interface Todo {
  id: number;
  text: string;
  created_at: string;
  updated_at: string;
  isEditing?: boolean;
}
