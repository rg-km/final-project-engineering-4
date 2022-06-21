import React from 'react';
import { Link } from 'react-router-dom';
import { Box, Heading, Stack, Text } from '@chakra-ui/react';
import AuthBackground from './AuthBackground';

function AuthFormContainer({ title, children }) {
  return (
    <AuthBackground>
      <Box
        bg={'white'}
        maxWidth={'2xl'}
        width={'full'}
        mx={{ base: 4, md: 'auto' }}
        my={'auto'}
        px={{ base: 10, md: 16 }}
        py={16}
        borderRadius={'md'}
        boxShadow={['md']}>
        <Stack spacing={12}>
          <Stack>
            <Heading textAlign={'center'} fontSize={['xl', '2xl', '3xl']}>
              {title}
            </Heading>
          </Stack>

          <Stack>{children}</Stack>

          <Stack>
            <Text textAlign={'center'} fontSize={'sm'}>
              Belum punya akun?{' '}
              <Link to={'/register'} style={{ textDecoration: 'underline' }}>
                daftar
              </Link>
            </Text>
          </Stack>
        </Stack>
      </Box>
    </AuthBackground>
  );
}

export default AuthFormContainer;
