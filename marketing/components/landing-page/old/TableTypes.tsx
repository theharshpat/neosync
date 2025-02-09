import Image from 'next/image';
import { ReactElement } from 'react';

export default function TableTypes(): ReactElement {
  const tableTypes = [
    {
      name: 'Tabular',
      description:
        'Generate statistically consistent synthetic data for a single data frame or table.',
      image: (
        <Image
          src="https://assets.nucleuscloud.com/neosync/marketingsite/tabularcropped.png"
          alt="pre"
          width="436"
          height="416"
          className="w-[240px] h-auto"
        />
      ),
    },
    {
      name: 'Relational',
      description:
        'Generate statistically consistent synthetic data for a relational database while maintaining referential integrity.',
      image: (
        <Image
          src="https://assets.nucleuscloud.com/neosync/marketingsite/relationaltablesgrid.png"
          alt="pre"
          width="436"
          height="416"
          className="w-[600px] h-auto"
        />
      ),
    },
  ];
  return (
    <div className="bg-[#F5F5F5] pt-20">
      <div className="flex flex-col  pt-5 lg:pt-40 bg-[#F5F5F5] px-5 sm:px-10 md:px-20 lg:px-80 max-w-[1800px] mx-auto ">
        <div className="text-gray-900 font-semibold text-2xl lg:text-5xl font-satoshi z-10  ">
          Generate Data for Any Schema
        </div>
        <div className="flex flex-col lg:flex-row space-y-4 lg:space-y-0 items-center justify-between gap-x-20 pt-10">
          {tableTypes.map((item) => (
            <div
              key={item.name}
              className="justify-between  border border-gray-400 bg-white shadow-lg rounded-xl p-6 lg:w-[800px] lg:h-[400px]"
            >
              <div className="justify-center flex">{item.image}</div>
              <div>
                <div className="text-xl text-gray-800 font-satoshi font-semibold pt-10">
                  {item.name}
                </div>
                <div className=" text-gray-600 font-satoshi pt-4 ">
                  {item.description}
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
