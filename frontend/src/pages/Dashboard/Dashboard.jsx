import {
  Button,
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
import { useAuthStore } from '@store/auth.store';
import { useState } from 'react';
import ClassItem from './components/ClassItem';
import { DATA_KELAS } from './data/kelas.data';

const initialInput = {
  nama_kelas: '',
  keterangan: '',
};

const Dashboard = () => {
  const { role } = useAuthStore((state) => state);
  const studentClasses = DATA_KELAS;
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [input, setInput] = useState(initialInput);

  const handleSubmit = () => {
    console.log(input);
  };

  return (
    <Stack py={'4'} spacing={'4'}>
      <Stack direction={'row'} justifyContent={'space-between'} alignItems={'center'}>
        <Heading fontSize={{ base: 'xs', md: 'md' }} mb={{ base: 2 }}>
          Daftar Kelas Anda
        </Heading>
        <Button onClick={onOpen}>Tambah Kelas Baru</Button>
      </Stack>

      <SimpleGrid columns={{ base: 2, md: 4 }} spacing={{ base: 2, md: 4 }}>
        {studentClasses.length > 0 &&
          studentClasses.map((kelas) => {
            return <ClassItem key={kelas.kode_kelas} to={PATH.DASHBOARD} kelas={kelas} />;
          })}
      </SimpleGrid>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader fontSize={{ base: 'md' }}>Tambah Kelas Baru</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            {role === 'guru' ? (
              <>
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
              </>
            ) : (
              <>
                <InputText type={'text'} label={'Kode Kelas'} required />
              </>
            )}
          </ModalBody>
          <ModalFooter>
            <Button colorScheme="blue" onClick={handleSubmit} size={{ base: 'sm' }}>
              Tambah Baru
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Stack>
  );
};

export default Dashboard;
