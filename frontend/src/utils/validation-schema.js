import * as yup from 'yup';
const { REGEX_PATTERN } = require('./regex-pattern');

const registerSchema = yup.object().shape({
  nama: yup.string().required('Nama harus diisi'),
  email: yup.string().email('Email tidak valid').required('Email harus diisi'),
  no_hp: yup
    .string()
    .matches(REGEX_PATTERN.number, 'No HP tidak valid')
    .min(10, 'No HP tidak valid')
    .required('No HP harus diisi'),
  password: yup.string().min(8, 'Password minimal 8 karakter').required('Password harus diisi'),
  confirm_password: yup
    .string()
    .oneOf([yup.ref('password'), null], 'Password tidak sama')
    .required('Konfirmasi password harus diisi'),
});

export const VALIDATION_SCHEMA = { registerSchema };
