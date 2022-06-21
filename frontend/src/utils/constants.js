export const USER_ROLES = {
  TEACHER: {
    value: 'teacher',
    title: 'Teacher',
    description: 'You are a leading educator in the nation.',
    icon: require('@images/icons/teaching.png'),
  },
  PARENT: {
    value: 'parent',
    title: 'Parent',
    description: 'They are wonderful as they will ever be!',
    icon: require('@images/icons/parents.png'),
  },
  STUDENT: {
    value: 'student',
    title: 'Student',
    description: 'in-built desire to learn and grow.',
    icon: require('@images/icons/student.png'),
  },
};

export const USER_ROLES_ARRAY = Object.values(USER_ROLES);
