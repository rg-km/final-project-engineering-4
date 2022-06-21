import { Link } from 'react-router-dom';
import { Container, Heading, Image, Stack, Text, Box } from '@chakra-ui/react';
import { APP, USER_ROLES_ARRAY } from '@utils/constants';

function Home() {
  return (
    <div>
      <Box
        bgImage={require('@images/picture/bg_auth.jpg')}
        height={'100vh'}
        display={'flex'}
        bgRepeat={'no-repeat'}
        bgSize={'cover'}
        objectFit={'cover'}>
        <Container maxWidth={'8xl'} py={'20'}>
          <Stack
            height={'full'}
            alignItems={'center'}
            justifyContent={'center'}
            textAlign={'center'}
            spacing={'10'}>
            <Stack spacing={'4'}>
              <Heading color={'primary.500'} fontSize={['3xl', '4xl', '5xl']}>
                {APP.name} {APP.version}
              </Heading>
              <Text color={'gray.50'}>Empowering education in Indonesia through technology</Text>
            </Stack>

            <Stack rounded={'md'} padding={[5, 8, 10]} bg={'gray.100'} alignItems={'center'} spacing={8}>
              <Text>Masuk Sebagai :</Text>

              <Stack direction={{ base: 'column', md: 'row' }} gap={{ base: 1, md: 3 }} justifyContent={'center'}>
                {USER_ROLES_ARRAY.map((role) => (
                  <Stack
                    as={Link}
                    key={role.value}
                    to={role.value}
                    bg={'white'}
                    rounded={'md'}
                    shadow={'sm'}
                    direction={{ base: 'row', md: 'column' }}
                    flex={1}
                    padding={{ base: 5, md: 6 }}
                    transition={'all 200ms ease'}
                    alignItems={'center'}
                    border={'1px'}
                    borderColor={'transparent'}
                    spacing={{ base: 6, md: 3 }}
                    _hover={{ borderColor: 'primary.500' }}>
                    <Stack width={['20', '16', null]} bg={'primary.100'} rounded={'full'}>
                      <Image maxWidth={'full'} src={role.icon} />
                    </Stack>
                    <Stack textAlign={{ base: 'left', md: 'center' }}>
                      <Text fontWeight={'medium'}>{role.title}</Text>
                      <Text color={'gray.500'} fontSize={'sm'}>
                        {role.description}
                      </Text>
                    </Stack>
                  </Stack>
                ))}
              </Stack>
            </Stack>
          </Stack>

          <Stack position={'fixed'} bottom={0} left={0} width={'full'} py={3}>
            <Text textAlign={'center'} color={'white'} fontSize={'xs'}>
              &copy; copyright 2022 Engineering 4 Final Project Ruangguru CAMP
            </Text>
          </Stack>
        </Container>
      </Box>
    </div>
  );
}

export default Home;