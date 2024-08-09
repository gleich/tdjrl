const things = Application('Things3');
const todos = things.lists.byId('TMLogbookListSource').toDos();

const today = new Date();
today.setHours(0, 0, 0, 0);
const yesterday = new Date(today);
yesterday.setDate(today.getDate() - 2);

const todayTimestamp = today.getTime();
const yesterdayTimestamp = yesterday.getTime();

const completedTask = todos
  .filter((todo) => {
    const completionDate = todo.completionDate();
    if (!completionDate) return false;
    const completedTimestamp = new Date(completionDate).getTime();
    return (
      completedTimestamp >= yesterdayTimestamp &&
      completedTimestamp < todayTimestamp
    );
  })
  .map((todo) => {
    const project = todo.project();
    const area = todo.area();

    return {
      id: todo.id(),
      name: todo.name(),
      status: todo.status(),
      notes: todo.notes(),
      tags: todo.tagNames(),
      due_date: todo.dueDate() && todo.dueDate().toISOString(),
      completion_date:
        todo.completionDate() && todo.completionDate().toISOString(),
      project: project && {
        id: project.id(),
        name: project.name(),
        tags: project.tagNames(),
        area: project.area() && {
          id: project.area().id(),
          name: project.area().name(),
          tags: project.area().tagNames(),
        },
      },
      area: area && {
        id: area.id(),
        name: area.name(),
        tags: area.tagNames(),
      },
    };
  });
console.log(JSON.stringify(completedTask));
