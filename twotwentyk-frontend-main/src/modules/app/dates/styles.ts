import { styled } from "styled-components";

export const FilterSectionWrapper = styled.div`
  p {
    font-weight: 600;
    font-size: 16px;
    line-height: 19px;
    margin-bottom: 12px;
  }
  margin-bottom: 20px;
`;

export const FilterGroupWrapper = styled.div`
  display: flex;
  align-items: center;
  & > :not(:first-child) {
    margin-left: 18px;
  }
  @media screen and (max-width: 1024px) {
    display: grid;
    /* grid-area: ; */
    grid-template-columns: repeat(3, 1fr);
    grid-gap: 20px 24px;
    & > :not(:first-child) {
      margin-left: 0px;
    }
  }
  @media screen and (max-width: 650px) {
    grid-template-columns: repeat(2, 1fr);
  }
  @media screen and (max-width: 425px) {
    grid-template-columns: 1fr;
  }
`;

export const SortButton = styled.div`
  display: flex;
  cursor: pointer;
  align-items: center;
  justify-content: center;
  /* width: 51px; */
  max-width: 40px;
  min-width: 40px !important;
  width: 100%;
  height: 40px;
  border-radius: 8px;
  background: #0e4fbc;
  box-shadow: 1px 1px 10px 0px rgba(0, 0, 0, 0.2);
`;

export const ViewOfferInfoWrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  b {
    font-size: 14px;
    font-weight: 600;
    line-height: 24px;
  }
  p {
    margin-top: 12px;
    font-size: 12px;
    font-weight: 400;
    line-height: 24px;
  }
  h1 {
    font-size: 22px;
    font-weight: 700;
  }
`;

export const OfferPreviewCard = styled.div<{ bg: string }>`
  box-shadow: 0px 0px 15.3925px rgba(0, 0, 0, 0.05);
  width: 250px;
  height: 300px;
  margin: auto;
  background: ${({ bg }) =>
    bg
      ? `url(${bg}) no-repeat, #CDCDCD `
      : "linear-gradient(0deg, #CDCDCD, #CDCDCD), #FFFFFF;"};
  background-blend-mode: luminosity;
  background-size: cover;
  background-position: center;
  position: relative;
  border-radius: 12px;
  contain: content;
  margin-bottom: 36px;
  border-radius: 0px 6.148px;
`;

export const CardRarity = styled.div`
  position: absolute;
  background: #615e5e;
  height: 30px;
  right: 0;
  top: 0;
  color: #fff;
  font-size: 14.346px;
  font-weight: 500;
`;

export const ButtonGroup = styled.div``;

export const OfferButton = styled.div`
  margin-top: 24px;
  border-radius: 8px;
  height: 45px;
  background: #000;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #fff;
  font-size: 14px;
  font-weight: 700;
  width: 100%;
  &.deny {
    background: #e31a4a;
  }
`;

export const CardGridWrapper = styled.div`
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  grid-gap: 24px;
  @media screen and (max-width: 1024px) {
    grid-template-columns: repeat(3, 1fr);
  }
  @media screen and (max-width: 650px) {
    grid-template-columns: repeat(2, 1fr);
  }
  @media screen and (max-width: 375px) {
    grid-template-columns: 1fr;
  }
`;

export const ViewDateCardWrapper = styled.div<{ isview?: string }>`
  max-width: 405px;
  width: 100%;
  z-index: 100;
  height: calc(100% - 125px);
  background-color: #fff;
  position: fixed;
  top: 125px;
  right: ${({ isview }) => (isview ? "0" : "-405px")};

  h2 {
    font-weight: 500;
    font-size: 22px;
    line-height: 27px;
    margin-bottom: 32px;
    width: 100%;
  }
  display: flex;
  align-items: center;
  flex-direction: column;
  @media screen and (max-width: 1024px) {
    top: 76px;
    height: calc(100% - 76px);
  }
`;

export const PreviewCardWrapper = styled.div`
  display: flex;
  width: 100%;
  justify-content: center;
`;

export const ViewDateCardContainer = styled.div`
  position: relative;
  overflow-y: auto;
  padding: 33px 52px;
  width: 100%;
  .sell-confirm-button {
    height: 45px;
    margin-top: 30px;
  }
`;

export const PropertiesWrapper = styled.div`
  margin-top: 16px;
  width: 100%;
`;

export const PropertiesHeader = styled.div<{ noborder?: boolean | string }>`
  padding-bottom: 10px;
  cursor: pointer;
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  color: #0e4fbc;
  border-bottom: ${({ noborder }) =>
    noborder ? "none" : "1px solid #00000022"};
  span {
    font-weight: 700;
    font-size: 14px;
    line-height: 17px;
  }
`;

export const PropertiesContent = styled.div`
  margin-top: 10px;
  & > :not(:first-child) {
    /* margin-top: 8px; */
  }
  &:not(:last-child) {
    border-bottom: 0.5px solid #00000022;
    padding-bottom: 10px;
  }
`;

export const PropertyItem = styled.div<{ nfttrigger?: boolean | string }>`
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 12px;
  height: 45px;
  line-height: 15px;
  text-transform: capitalize;
  padding: ${({ nfttrigger }) => (nfttrigger ? "0 10px" : "none")};
  p {
    font-weight: 600;
  }
  span {
    font-weight: 300;
  }
`;

export const PriceItem = styled.div`
  font-weight: 600;
  font-size: 12px;
  height: 33px;
  line-height: 15px;
  display: flex;
  justify-content: space-between;
  span.weak {
    font-weight: 400;
  }
`;

export const CloseButton = styled.div`
  display: flex;
  position: absolute;
  cursor: pointer;
  justify-content: center;
  align-items: center;
  font-size: 30px;
  line-height: 30px;
  right: 30px;
  top: 20px;
`;

export const SetPriceWrapper = styled.div`
  margin-top: 21px;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  padding: 20px 0;
  .price {
    display: flex;
    justify-content: space-between;
    width: 100%;
    align-items: center;
  }
  & > * {
    position: relative;
    z-index: 1;
  }
  & > :not(:last-child) {
    /* margin-bottom: 18px; */
  }
  & > p {
    font-weight: 400;
    width: 100%;
    font-size: 14px;
    line-height: 24px;
  }
  &::after {
    content: "";
    position: absolute;
    z-index: 0;
    top: 0;
    bottom: 0;
    left: -52px;
    right: -52px;
    background: #eff8ff;
  }
  h5 {
    font-weight: 400;
    font-size: 12px;
    line-height: 15px;
    text-transform: uppercase;
    color: #878686;
    margin: 0;
    /* width: 100%; */
  }
  h4 {
    /* width: 100%; */
    font-weight: 500;
    font-size: 18px;
    margin: 4px 0 0;
    line-height: 24px;
  }
  .sell-confirm-button {
    margin: 13px 0 0;
  }
`;
