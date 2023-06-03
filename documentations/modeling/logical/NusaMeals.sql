-- MySQL Script generated by MySQL Workbench
-- Tue May 16 18:44:43 2023
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema NusaMeals
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema NusaMeals
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `NusaMeals` DEFAULT CHARACTER SET utf8 ;
USE `NusaMeals` ;

-- -----------------------------------------------------
-- Table `NusaMeals`.`menu`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`menu` (
  `id_menu` INT NOT NULL AUTO_INCREMENT,
  `id_category` INT NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `price` DECIMAL NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id_menu`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`category`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`category` (
  `id_category` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_category`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`food`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`food` (
  `id_food` INT NOT NULL AUTO_INCREMENT,
  `id_category` INT NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `category` VARCHAR(45) NOT NULL,
  `country` VARCHAR(45) NOT NULL,
  `total_calorie` VARCHAR(45) NOT NULL,
  `description` TEXT NOT NULL,
  `ingredient` TEXT NOT NULL,
  `photos` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_food`),
  INDEX `fk_food_category_idx` (`id_category` ASC),
  CONSTRAINT `fk_food_category`
    FOREIGN KEY (`id_category`)
    REFERENCES `NusaMeals`.`category` (`id_category`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`drink`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`drink` (
  `id_drink` INT NOT NULL AUTO_INCREMENT,
  `id_category` INT NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `category` VARCHAR(45) NOT NULL,
  `country` VARCHAR(45) NOT NULL,
  `total_calorie` VARCHAR(45) NOT NULL,
  `description` TEXT NOT NULL,
  `ingredient` TEXT NOT NULL,
  `photos` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_drink`),
  INDEX `fk_drink_category_idx` (`id_category` ASC),
  CONSTRAINT `fk_drink_category`
    FOREIGN KEY (`id_category`)
    REFERENCES `NusaMeals`.`category` (`id_category`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`payment_type`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`payment_type` (
  `id_payment_type` INT NOT NULL AUTO_INCREMENT,
  `payment_type_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_payment_type`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`level`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`level` (
  `id_level` INT NOT NULL AUTO_INCREMENT,
  `level_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_level`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`user` (
  `id_user` INT NOT NULL AUTO_INCREMENT,
  `id_level` INT NOT NULL,
  `username` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_user`),
  INDEX `fk_user_level_idx` (`id_level` ASC),
  CONSTRAINT `fk_user_level`
    FOREIGN KEY (`id_level`)
    REFERENCES `NusaMeals`.`level` (`id_level`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`customer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`customer` (
  `id_customer` INT NOT NULL AUTO_INCREMENT,
  `id_user` INT NOT NULL,
  `full_name` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `gender` VARCHAR(45) NOT NULL,
  `phone_number` VARCHAR(45) NOT NULL,
  `address` TEXT NOT NULL,
  `registration_date` DATETIME NOT NULL,
  PRIMARY KEY (`id_customer`),
  INDEX `fk_customer_user_idx` (`id_user` ASC),
  CONSTRAINT `fk_customer_user`
    FOREIGN KEY (`id_user`)
    REFERENCES `NusaMeals`.`user` (`id_user`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`order`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`order` (
  `id_order` INT NOT NULL AUTO_INCREMENT,
  `id_customer` INT NOT NULL,
  `id_menu` INT NOT NULL,
  `id_payment_type` INT NOT NULL,
  `customer_name` VARCHAR(45) NOT NULL,
  `customer_phone` VARCHAR(45) NOT NULL,
  `type` VARCHAR(45) NOT NULL,
  `order_status` VARCHAR(45) NOT NULL,
  `order_date` DATETIME NOT NULL,
  `payment_status` VARCHAR(45) NOT NULL,
  `total_price` DECIMAL NOT NULL,
  `action` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_order`),
  INDEX `fk_order_payment_type_idx` (`id_payment_type` ASC),
  INDEX `fk_order_menu_idx` (`id_menu` ASC),
  INDEX `fk_order_customer_idx` (`id_customer` ASC),
  CONSTRAINT `fk_order_payment_type`
    FOREIGN KEY (`id_payment_type`)
    REFERENCES `NusaMeals`.`payment_type` (`id_payment_type`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_order_menu`
    FOREIGN KEY (`id_menu`)
    REFERENCES `NusaMeals`.`menu` (`id_menu`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_order_customer`
    FOREIGN KEY (`id_customer`)
    REFERENCES `NusaMeals`.`customer` (`id_customer`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`reservation`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`reservation` (
  `id_reservation` INT NOT NULL AUTO_INCREMENT,
  `customer_name` VARCHAR(45) NOT NULL,
  `phone_number` VARCHAR(45) NOT NULL,
  `date` DATE NOT NULL,
  `time_in` TIME NOT NULL,
  `time_out` TIME NOT NULL,
  `action` VARCHAR(45) NULL,
  PRIMARY KEY (`id_reservation`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`list_table`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`list_table` (
  `id_list_table` INT NOT NULL AUTO_INCREMENT,
  `number_of_table` INT NOT NULL,
  `seat` INT NOT NULL,
  `type` VARCHAR(45) NOT NULL,
  `status` VARCHAR(45) NOT NULL,
  `location` VARCHAR(45) NOT NULL,
  `photos` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_list_table`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`report`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`report` (
  `id_report` INT NOT NULL AUTO_INCREMENT,
  `cash` DECIMAL NOT NULL,
  `debi_card` DECIMAL NOT NULL,
  `e_wallet` DECIMAL NOT NULL,
  `total_amount` DECIMAL NOT NULL,
  `date` DATE NOT NULL,
  PRIMARY KEY (`id_report`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`payment`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`payment` (
  `id_payment` INT NOT NULL AUTO_INCREMENT,
  `id_order` INT NOT NULL,
  `id_payment_type` INT NOT NULL,
  `status` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_payment`),
  INDEX `fk_payment_payment_type_idx` (`id_payment_type` ASC),
  INDEX `fk_payment_order_idx` (`id_order` ASC),
  CONSTRAINT `fk_payment_payment_type`
    FOREIGN KEY (`id_payment_type`)
    REFERENCES `NusaMeals`.`payment_type` (`id_payment_type`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_payment_order`
    FOREIGN KEY (`id_order`)
    REFERENCES `NusaMeals`.`order` (`id_order`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`card_type`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`card_type` (
  `id_card_type` INT NOT NULL AUTO_INCREMENT,
  `card_type_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_card_type`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`ewallet_type`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`ewallet_type` (
  `id_ewallet_type` INT NOT NULL AUTO_INCREMENT,
  `ewallet_type_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_ewallet_type`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`ewallet_account`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`ewallet_account` (
  `id_ewallet` INT NOT NULL,
  `id_ewallet_type` INT NOT NULL,
  `id_costumer` INT NOT NULL,
  `ewallet_type` VARCHAR(45) NOT NULL,
  `ewallet_account_name` VARCHAR(45) NOT NULL,
  `ewallet_account_number` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id_ewallet`),
  INDEX `fk_ewallet_account_ewallet_type_idx` (`id_ewallet_type` ASC),
  INDEX `fk_ewallet_account_customer_idx` (`id_costumer` ASC),
  CONSTRAINT `fk_ewallet_account_ewallet_type`
    FOREIGN KEY (`id_ewallet_type`)
    REFERENCES `NusaMeals`.`ewallet_type` (`id_ewallet_type`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_ewallet_account_customer1`
    FOREIGN KEY (`id_costumer`)
    REFERENCES `NusaMeals`.`customer` (`id_customer`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`card_account`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`card_account` (
  `id_card` INT NOT NULL AUTO_INCREMENT,
  `id_card_type` INT NOT NULL,
  `id_customer` INT NOT NULL,
  `card_holder_name` VARCHAR(45) NOT NULL,
  `card_number` VARCHAR(45) NOT NULL,
  `card_type` VARCHAR(45) NOT NULL,
  `exp_date` VARCHAR(45) NOT NULL,
  `cvv` INT NOT NULL,
  PRIMARY KEY (`id_card`),
  INDEX `fk_card_account_card_type_idx` (`id_card_type` ASC),
  INDEX `fk_card_account_customer_idx` (`id_customer` ASC),
  CONSTRAINT `fk_card_account_card_type`
    FOREIGN KEY (`id_card_type`)
    REFERENCES `NusaMeals`.`card_type` (`id_card_type`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_card_account_customer`
    FOREIGN KEY (`id_customer`)
    REFERENCES `NusaMeals`.`customer` (`id_customer`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `NusaMeals`.`order_cancel`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `NusaMeals`.`order_cancel` (
  `id_order_cancel` INT NOT NULL AUTO_INCREMENT,
  `id_order` INT NOT NULL,
  `id_customer` INT NOT NULL,
  `reason` TEXT NOT NULL,
  PRIMARY KEY (`id_order_cancel`),
  INDEX `fk_order_cancel_order_idx` (`id_order` ASC),
  INDEX `fk_order_cancel_customer_idx` (`id_customer` ASC),
  CONSTRAINT `fk_order_cancel_order`
    FOREIGN KEY (`id_order`)
    REFERENCES `NusaMeals`.`order` (`id_order`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_order_cancel_customer`
    FOREIGN KEY (`id_customer`)
    REFERENCES `NusaMeals`.`customer` (`id_customer`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;