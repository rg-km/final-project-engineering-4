import { Box, Divider, Heading, Stack } from '@chakra-ui/react';
import AuthBackground from './AuthBackground';

function AuthFormContainer({ title, children }) {
  return (
    <AuthBackground>
      <Box
        bg={'white'}
        position={'relative'}
        maxWidth={'xl'}
        width={'full'}
        mx={{ base: 4, md: 'auto' }}
        my={'auto'}
        px={{ base: 10, md: 12 }}
        py={12}
        borderRadius={'md'}
        boxShadow={['md']}>
        <Stack spacing={8}>
          <Stack>
            <Heading textAlign={'center'} fontSize={['lg', 'xl', '2xl']}>
              {title}
            </Heading>
          </Stack>

          <Divider />

          <Stack>{children}</Stack>
        </Stack>
      </Box>
    </AuthBackground>
  );
}

export default AuthFormContainer;
