import { Link } from 'react-router-dom';
import { Button, Stack, Text } from '@chakra-ui/react';
import AuthContainer from './components/AuthContainer';
import { PATH } from '@config/path';
import { USER_ROLES } from '@utils/constants';
import RegisterTeacher from './forms/RegisterTeacher';

const Register = ({ role }) => {
  return (
    <AuthContainer
      title={`Daftar sebagai ${role.title}`}
      subtitle="Lengkapi form di bawah dengan menggunakan data Anda yang valid">
      <Stack>
        {role.value === USER_ROLES.TEACHER.value && <RegisterTeacher />}

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
