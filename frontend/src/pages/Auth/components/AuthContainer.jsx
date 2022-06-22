import { Link } from 'react-router-dom';
import { Button, Container, Divider, Heading, Stack, Text } from '@chakra-ui/react';
import { PATH } from '@config/path';
import AuthBackground from './AuthBackground';

function AuthFormContainer({ title, subtitle, children }) {
  return (
    <AuthBackground>
      <Container as={Stack} justifyContent={'center'} maxWidth={'xl'} spacing={5}>
        <Stack
          bg={'white'}
          width={'full'}
          maxHeight={'80vh'}
          overflowY={'auto'}
          px={{ base: 10, md: 12 }}
          py={12}
          borderRadius={'md'}
          boxShadow={['md']}>
          <Stack spacing={8}>
            <Stack textAlign={'center'}>
              <Heading fontSize={['lg', 'xl', '2xl']}>{title}</Heading>
              {subtitle && <Text color={'gray.500'}>{subtitle}</Text>}
            </Stack>

            <Divider />

            <Stack>{children}</Stack>
          </Stack>
        </Stack>

        <Stack>
          <Button size={'sm'} as={Link} to={PATH.LOGIN} variant={'link'}>
            Kembali ke halaman utama
          </Button>
        </Stack>
      </Container>
    </AuthBackground>
  );
}

export default AuthFormContainer;
