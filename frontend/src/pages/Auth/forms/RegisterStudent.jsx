import { Button, Stack, useToast } from '@chakra-ui/react';
import React from 'react';
import InputPassword from 'src/components/InputPassword';
import InputText from 'src/components/InputText';
import { Form, Formik } from 'formik';
import { useNavigate } from 'react-router-dom';
import { useAuthStore } from '@store/auth.store';
import { USER_ROLES } from '@utils/constants';
import { VALIDATION_SCHEMA } from '@utils/validation-schema';

const INITIAL_VALUES = {
  nama: '',
  email: '',
  jenis_kelamin: '',
  no_hp: '',
  alamat: '',
  tempat_lahir: '',
  tanggal_lahir: '',
  agama: '',
  file_name: '',
  password: '',
  confirm_password: '',
};

const RegisterStudent = () => {
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
    tempat_lahir,
    tanggal_lahir,
    agama,
    file_name,
    password,
    confirm_password,
  }) => {
    const data = { nama, email, jenis_kelamin, no_hp, alamat, tempat_lahir, tanggal_lahir, agama, file_name, password, confirm_password };

    handleRegister(USER_ROLES.STUDENT.value, data, (success, message) => {
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
      validationSchema={VALIDATION_SCHEMA.registerStudentSchema}>
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
                name="tempat_lahir"
                label="Tempat Lahir"
                value={values.tempat_lahir}
                meta={getFieldMeta('tempat_lahir')}
                onChange={handleChange}
              />
              <InputText
                name="tanggal_lahir"
                label="Tanggal Lahir"
                value={values.tanggal_lahir}
                meta={getFieldMeta('tanggal_lahir')}
                onChange={handleChange}
              />
              <InputText
                name="agama"
                label="Agama"
                value={values.agama}
                meta={getFieldMeta('agama')}
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

export default RegisterStudent;
