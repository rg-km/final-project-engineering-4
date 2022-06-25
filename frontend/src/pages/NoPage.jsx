import { Box, Button, Image, Text } from "@chakra-ui/react";
import nopage from "@images/icons/nopage.svg";
import { Link } from "react-router-dom";

const NoPage = () => {
    return <Box h={'100vh'}>
        <Box m={"auto"} w={{base: "80%"}}>
            <Image src={nopage} alt="404" w={{base: "100%", md: "40%"}} m={"auto"} />
            <Box textAlign={"center"}>
                <Text mb={{base: 2}}>Halaman tidak ada</Text>
                <Button as={Link} to={"/dashboard"}>Kembali ke dashboard</Button>
            </Box>
        </Box>
    </Box>
}

export default NoPage;