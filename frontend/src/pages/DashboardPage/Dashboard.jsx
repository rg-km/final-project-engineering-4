import {
  Box,
  Button,
  Container,
  Heading,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  SimpleGrid,
  Stack,
  useDisclosure,
} from '@chakra-ui/react';
import InputText from '@components/InputText';
import { PATH } from '@config/path';
import { useState } from 'react';
import ClassItem from './components/ClassItem';
import Navbar from './components/Navbar';

const DATA = [
  {
    id_class: 1,
    nama_kelas: '9F',
    kode_kelas: 'qwertyasdf',
    keterangan: 'Kelas dengan wali kelas yang bernama Bu Nurul Huda',
  },
  {
    id_class: 2,
    nama_kelas: '9E',
    kode_kelas: 'qwertyuiop',
    keterangan: 'Kelas dengan wali kelas yang bernama Bu Sum',
  },
  {
    id_class: 3,
    nama_kelas: '9D',
    kode_kelas: 'zxcvasdf',
    keterangan: 'Kelas dengan wali kelas yang bernama Pak Rizky',
  },
];

const initialInput = {
  nama_kelas: '',
  keterangan: '',
};

const Dashboard = () => {
  const studentClasses = DATA;
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [input, setInput] = useState(initialInput);

  const handleSubmit = () => {
    console.log(input);
  };

  return (
    <Box bg={'gray.100'} minH={'100vh'}>
      <Navbar />

      <Container maxWidth={'8xl'}>
        <Stack py={'4'} spacing={'4'}>
          <Stack direction={'row'} justifyContent={'space-between'} alignItems={'center'}>
            <Heading fontSize={{ base: 'xs', md: 'md' }} mb={{ base: 2 }}>
              Daftar Kelas Anda
            </Heading>
            <Button onClick={onOpen}>Tambah Kelas Baru</Button>
          </Stack>

          <SimpleGrid columns={{ base: 2, md: 4 }} spacing={{ base: 2, md: 4 }}>
            {studentClasses.length > 0 &&
              studentClasses.map((kelas, index) => {
                return <ClassItem key={index} to={PATH.DASHBOARD} kelas={kelas} />;
              })}
          </SimpleGrid>

          <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay />
            <ModalContent>
              <ModalHeader fontSize={{ base: 'md' }}>Tambah Kelas Baru</ModalHeader>
              <ModalCloseButton />
              <ModalBody>
                <InputText
                  type={'text'}
                  label={'Nama Kelas'}
                  value={input.nama_kelas}
                  onChange={(e) => setInput({ ...input, nama_kelas: e.target.value })}
                  required
                />
                <InputText
                  type={'text'}
                  label={'Keterangan'}
                  value={input.keterangan}
                  onChange={(e) => setInput({ ...input, keterangan: e.target.value })}
                />
              </ModalBody>
              <ModalFooter>
                <Button colorScheme="blue" onClick={handleSubmit} size={{ base: 'sm' }}>
                  Tambah Baru
                </Button>
              </ModalFooter>
            </ModalContent>
          </Modal>
        </Stack>
      </Container>
    </Box>
  );
};

export default Dashboard;
