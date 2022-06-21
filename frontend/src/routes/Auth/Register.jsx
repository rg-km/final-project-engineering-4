import { Link } from 'react-router-dom';
import { Button, Stack, Text } from '@chakra-ui/react';
import { Formik, Form } from 'formik';
import { PATH } from '@routes/path';
import AuthContainer from './components/AuthContainer';
import InputText from '@components/InputText';
import InputPassword from '@components/InputPassword';

const INITIAL_VALUES = {
  nama: '',
  username: '',
  email: '',
  jenis_kelamin: '',
  pendidikan: '',
  agama: '',
  no_hp: '',
  alamat: '',
  password: '',
  confirm_password: '',
};

const Register = ({ role }) => {
  const handleSubmit = (values) => {
    console.log({ values });
  };

  return (
    <AuthContainer
      title={`Daftar sebagai ${role.title}`}
      subtitle="Lengkapi form di bawah dengan menggunakan data Anda yang valid">
      <Formik initialValues={INITIAL_VALUES} onSubmit={handleSubmit}>
        {({ values, handleChange }) => (
          <Form>
            <Stack spacing={8}>
              <Stack spacing={3}>
                <InputText name="nama" label="Nama" value={values.nama} onChange={handleChange} />
                <InputText name="username" label="Username" value={values.username} onChange={handleChange} />
                <InputText name="email" label="Email" value={values.email} onChange={handleChange} />
                <InputText
                  name="jenis_kelamin"
                  label="Jenis Kelamin"
                  value={values.jenis_kelamin}
                  onChange={handleChange}
                />
                <InputText
                  name="pendidikan"
                  label="Pendidikan"
                  value={values.pendidikan}
                  onChange={handleChange}
                />
                <InputText name="agama" label="Agama" value={values.agama} onChange={handleChange} />
                <InputText name="no_hp" label="No HP" value={values.no_hp} onChange={handleChange} />
                <InputText name="alamat" label="Alamat" value={values.alamat} onChange={handleChange} />
                <InputPassword name="password" label="Password" value={values.password} onChange={handleChange} />
                <InputPassword
                  name="confirm_password"
                  label="Konfirmasi Password"
                  value={values.confirm_password}
                  onChange={handleChange}
                />
              </Stack>

              <Button type="submit">Daftar</Button>
            </Stack>
          </Form>
        )}
      </Formik>

      <Stack>
        <Text textAlign={'center'} fontSize={'sm'}>
          Sudah punya akun?{' '}
          <Button as={Link} size={'sm'} variant={'link'} to={`${PATH.LOGIN}/${role.value}`}>
            Masuk sebagai {role.value}
          </Button>
        </Text>
      </Stack>
    </AuthContainer>
  );
};

export default Register;
