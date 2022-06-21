import { Box, Button, Flex, Heading, Text } from '@chakra-ui/react';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import AppAlert from 'src/components/AppAlert';
import AuthBackground from '@routes/Auth/components/AuthBackground';
import InputPassword from 'src/components/InputPassword';
import InputText from 'src/components/InputText';

const LoginParent = () => {
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
    <AuthBackground>
      <Box
        maxW={'500px'}
        overflow={'auto'}
        bg={'white'}
        maxH={'80vh'}
        w={['90%', '50%']}
        m={'auto'}
        p={[4, 6, 10, 20]}
        borderRadius={['xl', '2xl']}
        boxShadow={['2xl']}>
        <Heading textAlign={'center'} fontSize={['sm', 'md', 'lg']} mb={[3, 4, 8]}>
          Masuk sebagai Orang Tua
        </Heading>

        <AppAlert alert={alert} setAlert={setAlert} />

        <form onSubmit={onSubmit}>
          <InputText name="email" label="Email" value={input.email} handleChange={handleChange} isRequired />
          <InputPassword
            name="password"
            label="Password"
            value={input.password}
            handleChange={handleChange}
            isRequired
          />

          <Flex justify={'center'} gap={[2]} mb={[2, 4]}>
            <Link to={'/login'}>
              <Button size={['sm', 'md']}>Kembali</Button>
            </Link>
            <Button type="submit" isLoading={isLoading} colorScheme="blue" size={['sm', 'md']}>
              Login
            </Button>
          </Flex>
        </form>

        <Text textAlign={'center'} fontSize={['xs', 'sm']}>
          Belum punya akun?{' '}
          <Link to={'/register'} style={{ textDecoration: 'underline' }}>
            daftar
          </Link>
        </Text>
      </Box>
    </AuthBackground>
  );
};

export default LoginParent;
