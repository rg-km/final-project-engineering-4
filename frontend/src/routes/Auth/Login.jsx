import { Button, Stack, Text } from '@chakra-ui/react';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import AppAlert from '@components/AppAlert';
import InputPassword from '@components/InputPassword';
import InputText from '@components/InputText';
import AuthFormContainer from './components/AuthFormContainer';
import { PATH } from '@routes/path';

const Login = ({ role }) => {
  const initilInput = {
    email: '',
    password: '',
  };
  const [isLoading, setIsLoading] = useState(false);
  const [input, setInput] = useState(initilInput);
  const [alert, setAlert] = useState({
    status: false,
    type: 'warning',
    message: '',
  });

  const handleChange = (e) => {
    setInput({
      ...input,
      [e.target.name]: e.target.value,
    });
  };

  const onSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    console.log(input);
    setIsLoading(false);
  };

  return (
    <AuthFormContainer title={`Masuk sebagai ${role.title}`}>
      <AppAlert alert={alert} setAlert={setAlert} />

      <form onSubmit={onSubmit}>
        <Stack spacing={8}>
          <Stack spacing={3}>
            <InputText name="email" label="Email" value={input.email} handleChange={handleChange} isRequired />
            <InputPassword
              name="password"
              label="Password"
              value={input.password}
              handleChange={handleChange}
              isRequired
            />
          </Stack>

          <Stack direction={'row'}>
            <Button type="submit" variant={'outline'} isLoading={isLoading} flex={1}>
              Kembali
            </Button>
            <Button type="submit" isLoading={isLoading} flex={1}>
              Login
            </Button>
          </Stack>

          <Stack>
            <Text textAlign={'center'} fontSize={'sm'}>
              Belum punya akun?{' '}
              <Button as={Link} size={'sm'} variant={'link'} to={`${PATH.REGISTER}/${role.value}`}>
                Daftar
              </Button>
            </Text>
          </Stack>
        </Stack>
      </form>
    </AuthFormContainer>
  );
};

export default Login;
