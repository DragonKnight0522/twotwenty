import React, { useEffect } from "react";
import { ModalProps } from "../../types";
import {
  CloseButton,
  ModalContainer,
  ModalOverlay,
  ModalWrapper,
} from "./styles";
import { ModalHeader } from "./ModalHeader";

export const Modal: React.FC<ModalProps> = ({
  onClose,
  open,
  children,
  width,
  modalHeader,
  paddingClass
}) => {
  useEffect(() => {
    if (open) {
      document.documentElement.style.overflow = "hidden";
    } else {
      document.documentElement.style.overflow = "auto";
    }
  }, [open]);
  return (
    <ModalWrapper open={open}>
      <ModalContainer width={width} className={paddingClass}>
        {modalHeader ? (
          <ModalHeader onClose={onClose} />
        ) : (
          <CloseButton onClick={onClose}>&times;</CloseButton>
        )}
        {children}
      </ModalContainer>
      <ModalOverlay onClick={onClose} />
    </ModalWrapper>
  );
};
