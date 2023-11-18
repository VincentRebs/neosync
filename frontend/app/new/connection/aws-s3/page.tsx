import OverviewContainer from '@/components/containers/OverviewContainer';
import PageHeader from '@/components/headers/PageHeader';
import { FaAws } from 'react-icons/fa';
import AwsS3Form from './AwsS3Form';

export default async function Postgres() {
  return (
    <OverviewContainer
      Header={
        <PageHeader
          header="AWS S3"
          description="Configure an AWS S3 data connection"
          pageHeaderContainerClassName="mx-64"
          leftIcon={<FaAws className="w-[40px] h-[40px]" />}
        />
      }
    >
      <AwsS3Form />
    </OverviewContainer>
  );
}
