import { Box, Flex, Heading, Text } from '@chakra-ui/react';
import React from 'react';
import { Link } from 'react-router-dom';
import RoleItem from 'src/components/RoleItem';
import teacher from '@images/icons/teaching.png';
import parent from '@images/icons/parents.png';
import student from '@images/icons/student.png';
import BgAuth from 'src/components/BgAuth';

const Register = () => {
  return (
    <BgAuth>
      <Box
        bg={'white'}
        w={['90%', '50%']}
        m={'auto'}
        p={[4, 6, 10, 20]}
        borderRadius={['xl', '2xl']}
        boxShadow={['2xl']}>
        <Heading textAlign={'center'} fontSize={['sm', 'md', 'lg']} mb={[3, 4, 8]}>
          Daftar sebagai
        </Heading>
        <Flex gap={[2, 4, 8]} justifyContent={'center'} mb={[2, 4]}>
          <RoleItem title={'Guru'} redirect={'/register/teacher'} img={teacher} bg={'yellow.200'} />
          <RoleItem
            title={'Orang Tua'}
            redirect={'/register/parent'}
            img={parent}
            bg={'blue.200'}
          />
          <RoleItem title={'Siswa'} redirect={'/register/student'} img={student} bg={'green.200'} />
        </Flex>
        <Text textAlign={'center'} fontSize={['xs', 'sm']}>
          Atau{' '}
          <Link to={'/login'} style={{ textDecoration: 'underline' }}>
            masuk
          </Link>
        </Text>
      </Box>
    </BgAuth>
  );
};

export default Register;
