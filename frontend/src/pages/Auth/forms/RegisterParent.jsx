import { Button, Stack, useToast } from '@chakra-ui/react';
import { Form, Formik } from 'formik';
import React from 'react';
import InputPassword from 'src/components/InputPassword';
import InputText from 'src/components/InputText';
import { USER_ROLES } from '@utils/constants';
import { useNavigate } from 'react-router-dom';
import { useAuthStore } from '@store/auth.store';
import { VALIDATION_SCHEMA } from '@utils/validation-schema';

const INITIAL_VALUES = {
  nama: '',
  email: '',
  jenis_kelamin: '',
  no_hp: '',
  alamat: '',
  email_siswa: '',
  password: '',
  confirm_password: '',
};

const RegisterParent = () => {
  const toast = useToast();
  const navigate = useNavigate();

  const { isLoading } = useAuthStore((state) => state); // data
  const { handleRegister } = useAuthStore((state) => state); // action

  const handleSubmit = ({
    nama,
    email,
    jenis_kelamin,
    no_hp,
    alamat,
    email_siswa,
    password,
    confirm_password,
  }) => {
    const data = { nama, email, jenis_kelamin, no_hp, alamat, email_siswa, password, confirm_password };

    handleRegister(USER_ROLES.PARENT.value, data, (success, message) => {
      toast({
        status: success ? 'success' : 'error',
        position: 'top-right',
        title: message,
        duration: 2500,
        isClosable: true,
      });

      navigate(PATH.LOGIN);
    });
  };

  return (
    <Formik
      initialValues={INITIAL_VALUES}
      onSubmit={handleSubmit}
      validationSchema={VALIDATION_SCHEMA.registerParentSchema}>
      {({ values, handleSubmit, getFieldMeta, handleChange }) => (
        <Form onSubmit={handleSubmit}>
          <Stack spacing={8}>
            <Stack spacing={3}>
              <InputText
                name="nama"
                label="Nama"
                value={values.nama}
                meta={getFieldMeta('nama')}
                onChange={handleChange}
              />
              <InputText
                name="email"
                label="Email"
                value={values.email}
                meta={getFieldMeta('email')}
                onChange={handleChange}
              />
              <InputText
                name="jenis_kelamin"
                label="Jenis Kelamin"
                meta={getFieldMeta('jenis_kelamin')}
                value={values.jenis_kelamin}
                onChange={handleChange}
              />
              <InputText
                name="no_hp"
                label="No HP"
                value={values.no_hp}
                meta={getFieldMeta('no_hp')}
                onChange={handleChange}
              />
              <InputText
                name="alamat"
                label="Alamat"
                value={values.alamat}
                meta={getFieldMeta('alamat')}
                onChange={handleChange}
              />
              <InputText
                name="email_siswa"
                label="Email Siswa"
                value={values.email_siswa}
                meta={getFieldMeta('email_siswa')}
                onChange={handleChange}
              />
              <InputPassword
                name="password"
                label="Password"
                value={values.password}
                meta={getFieldMeta('password')}
                onChange={handleChange}
              />
              <InputPassword
                name="confirm_password"
                label="Konfirmasi Password"
                value={values.confirm_password}
                meta={getFieldMeta('confirm_password')}
                onChange={handleChange}
              />
            </Stack>

            <Button type="submit" isLoading={isLoading}>
              Daftar
            </Button>
          </Stack>
        </Form>
      )}
    </Formik>
  );
};

export default RegisterParent;
