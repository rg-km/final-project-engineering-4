import { Link } from 'react-router-dom';
import { Form, Formik } from 'formik';
import { Button, Stack, Text, useToast } from '@chakra-ui/react';
import InputPassword from '@components/InputPassword';
import InputText from '@components/InputText';
import AuthContainer from './components/AuthContainer';
import { PATH } from '@config/path';
import { VALIDATION_SCHEMA } from '@utils/validation-schema';
import { useAuthStore } from '@store/auth.store';

const INITIAL_VALUES = {
  email: '',
  password: '',
};

const Login = ({ role }) => {
  const toast = useToast();

  const { isLoading } = useAuthStore((state) => state); // data
  const { handleLogin } = useAuthStore((state) => state); // action

  const handleSubmit = async (values) => {
    handleLogin(role, values.email, values.password, (success, message) => {
      toast({
        status: success ? 'success' : 'error',
        position: 'top-right',
        title: message,
        duration: 2500,
        isClosable: true,
      });
    });
  };

  return (
    <AuthContainer title={`Masuk sebagai ${role.title}`} subtitle="Masukkan email dan password Anda untuk masuk.">
      <Formik
        initialValues={INITIAL_VALUES}
        validationSchema={VALIDATION_SCHEMA.loginSchema}
        onSubmit={handleSubmit}>
        {({ values, handleChange, handleSubmit, getFieldMeta }) => (
          <Stack as={Form} onSubmit={handleSubmit} spacing={8}>
            <Stack spacing={3}>
              <InputText
                name="email"
                label="Email"
                value={values.email}
                onChange={handleChange}
                meta={getFieldMeta('email')}
                disabled={isLoading}
                required
              />
              <InputPassword
                name="password"
                label="Password"
                value={values.password}
                meta={getFieldMeta('password')}
                onChange={handleChange}
                disabled={isLoading}
                required
              />
            </Stack>

            <Stack direction={'row'}>
              <Button type="submit" flex={1} isLoading={isLoading}>
                Login
              </Button>
            </Stack>
          </Stack>
        )}
      </Formik>
      <Stack>
        <Text textAlign={'center'} fontSize={'sm'}>
          Belum punya akun?{' '}
          <Button as={Link} size={'sm'} variant={'link'} to={`${PATH.REGISTER}/${role.value}`}>
            Daftar sebagai {role.value}
          </Button>
        </Text>
      </Stack>
    </AuthContainer>
  );
};

export default Login;
