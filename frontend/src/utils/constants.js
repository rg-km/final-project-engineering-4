export const APP = {
  name: 'MasukKelas',
  version: '1.0.0',
  copyright: '© copyright 2022 Engineering 4 Final Project Ruangguru CAMP',
};

export const USER_ROLES = {
  TEACHER: {
    value: 'guru',
    title: 'Guru',
    description: 'You are a leading educator in the nation.',
    icon: require('@images/icons/teaching.png'),
  },
  PARENT: {
    value: 'parent',
    title: 'Orang Tua',
    description: 'They are wonderful as they will ever be!',
    icon: require('@images/icons/parents.png'),
  },
  STUDENT: {
    value: 'student',
    title: 'Siswa',
    description: 'in-built desire to learn and grow.',
    icon: require('@images/icons/student.png'),
  },
};

export const USER_ROLES_ARRAY = Object.values(USER_ROLES);
