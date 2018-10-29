package main

import (
	"errors"
)

const (
	HEALTH_PROP_2000_VERSION_MIN           int16   = 2000
	HEALTH_PROP_2000_BASIS_MONTHLY_MINIMUM int32   = 0
	HEALTH_PROP_2000_BASIS_ANNUAL_MAXIMUM  int32   = 0
	HEALTH_PROP_2000_FACTOR_COMPOUND       float32 = 0
	HEALTH_PROP_2000_INCOME_EMPLOY_MARGIN  int64   = 0
	HEALTH_PROP_2000_INCOME_AGREEM_MARGIN  int64   = 0

	HEALTH_PROP_2011_VERSION_MIN           int16   = 2011
	HEALTH_PROP_2011_BASIS_MONTHLY_MINIMUM int32   = 8000
	HEALTH_PROP_2011_BASIS_ANNUAL_MAXIMUM  int32   = 1781280
	HEALTH_PROP_2011_FACTOR_COMPOUND       float32 = 13.5
	HEALTH_PROP_2011_INCOME_EMPLOY_MARGIN  int64   = 2000
	HEALTH_PROP_2011_INCOME_AGREEM_MARGIN  int64   = 0

	HEALTH_PROP_2012_VERSION_MIN           int16   = 2012
	HEALTH_PROP_2012_BASIS_MONTHLY_MINIMUM int32   = HEALTH_PROP_2011_BASIS_MONTHLY_MINIMUM
	HEALTH_PROP_2012_BASIS_ANNUAL_MAXIMUM  int32   = 1809864
	HEALTH_PROP_2012_FACTOR_COMPOUND       float32 = HEALTH_PROP_2011_FACTOR_COMPOUND
	HEALTH_PROP_2012_INCOME_EMPLOY_MARGIN  int64   = 2500
	HEALTH_PROP_2012_INCOME_AGREEM_MARGIN  int64   = 10000

	HEALTH_PROP_2013_VERSION_MIN           int16   = 2013
	HEALTH_PROP_2013_BASIS_MONTHLY_MINIMUM int32   = 8500
	HEALTH_TO07_2013_BASIS_MONTHLY_MINIMUM int32   = HEALTH_PROP_2012_BASIS_MONTHLY_MINIMUM
	HEALTH_PROP_2013_BASIS_ANNUAL_MAXIMUM  int32   = 0
	HEALTH_PROP_2013_FACTOR_COMPOUND       float32 = 0
	HEALTH_PROP_2013_INCOME_EMPLOY_MARGIN  int64   = HEALTH_PROP_2012_INCOME_EMPLOY_MARGIN
	HEALTH_PROP_2013_INCOME_AGREEM_MARGIN  int64   = HEALTH_PROP_2012_INCOME_AGREEM_MARGIN

	HEALTH_PROP_2014_VERSION_MIN           int16   = 2014
	HEALTH_PROP_2014_BASIS_MONTHLY_MINIMUM int32   = HEALTH_PROP_2013_BASIS_MONTHLY_MINIMUM
	HEALTH_PROP_2014_BASIS_ANNUAL_MAXIMUM  int32   = HEALTH_PROP_2013_BASIS_ANNUAL_MAXIMUM
	HEALTH_PROP_2014_FACTOR_COMPOUND       float32 = HEALTH_PROP_2013_FACTOR_COMPOUND
	HEALTH_PROP_2014_INCOME_EMPLOY_MARGIN  int64   = HEALTH_PROP_2013_INCOME_EMPLOY_MARGIN
	HEALTH_PROP_2014_INCOME_AGREEM_MARGIN  int64   = HEALTH_PROP_2013_INCOME_AGREEM_MARGIN

	HEALTH_PROP_2015_VERSION_MIN           int16   = 2015
	HEALTH_PROP_2015_BASIS_MONTHLY_MINIMUM int32   = 9200
	HEALTH_PROP_2015_BASIS_ANNUAL_MAXIMUM  int32   = HEALTH_PROP_2014_BASIS_ANNUAL_MAXIMUM
	HEALTH_PROP_2015_FACTOR_COMPOUND       float32 = HEALTH_PROP_2014_FACTOR_COMPOUND
	HEALTH_PROP_2015_INCOME_EMPLOY_MARGIN  int64   = HEALTH_PROP_2014_INCOME_EMPLOY_MARGIN
	HEALTH_PROP_2015_INCOME_AGREEM_MARGIN  int64   = HEALTH_PROP_2014_INCOME_AGREEM_MARGIN

	HEALTH_PROP_2016_VERSION_MIN           int16   = 2016
	HEALTH_PROP_2016_BASIS_MONTHLY_MINIMUM int32   = 9900
	HEALTH_PROP_2016_BASIS_ANNUAL_MAXIMUM  int32   = HEALTH_PROP_2015_BASIS_ANNUAL_MAXIMUM
	HEALTH_PROP_2016_FACTOR_COMPOUND       float32 = HEALTH_PROP_2015_FACTOR_COMPOUND
	HEALTH_PROP_2016_INCOME_EMPLOY_MARGIN  int64   = HEALTH_PROP_2015_INCOME_EMPLOY_MARGIN
	HEALTH_PROP_2016_INCOME_AGREEM_MARGIN  int64   = HEALTH_PROP_2015_INCOME_AGREEM_MARGIN

	HEALTH_PROP_2017_VERSION_MIN           int16   = 2017
	HEALTH_PROP_2017_BASIS_MONTHLY_MINIMUM int32   = 11000
	HEALTH_PROP_2017_BASIS_ANNUAL_MAXIMUM  int32   = HEALTH_PROP_2016_BASIS_ANNUAL_MAXIMUM
	HEALTH_PROP_2017_FACTOR_COMPOUND       float32 = HEALTH_PROP_2016_FACTOR_COMPOUND
	HEALTH_PROP_2017_INCOME_EMPLOY_MARGIN  int64   = HEALTH_PROP_2016_INCOME_EMPLOY_MARGIN
	HEALTH_PROP_2017_INCOME_AGREEM_MARGIN  int64   = HEALTH_PROP_2016_INCOME_AGREEM_MARGIN

	HEALTH_PROP_2018_VERSION_MIN           int16   = 2018
	HEALTH_PROP_2018_BASIS_MONTHLY_MINIMUM int32   = 12200
	HEALTH_PROP_2018_BASIS_ANNUAL_MAXIMUM  int32   = HEALTH_PROP_2017_BASIS_ANNUAL_MAXIMUM
	HEALTH_PROP_2018_FACTOR_COMPOUND       float32 = HEALTH_PROP_2017_FACTOR_COMPOUND
	HEALTH_PROP_2018_INCOME_EMPLOY_MARGIN  int64   = HEALTH_PROP_2017_INCOME_EMPLOY_MARGIN
	HEALTH_PROP_2018_INCOME_AGREEM_MARGIN  int64   = HEALTH_PROP_2017_INCOME_AGREEM_MARGIN
)

type HealthGuides struct {
	InternalPeriod      Period
	BasisMonthlyMinimum int32
	BasisAnnualMaximum  int32
	FactorCompound      float32
	IncomeEmployMargin  int64
	IncomeAgreemMargin  int64
}

func NewHealthGuides(period Period, versionMin int16,
	basisMonthlyMinimum int32, basisAnnualMaximum int32,
	factorCompound float32, incomeEmployMargin int64, incomeAgreemMargin int64) (guides *HealthGuides, err error) {

	if period.Year() < versionMin || period.Year() > versionMin {
		return nil, errors.New("Period is not valid for this version")
	}
	guides = &HealthGuides{
		InternalPeriod:      period,
		BasisMonthlyMinimum: basisMonthlyMinimum,
		BasisAnnualMaximum:  basisAnnualMaximum,
		FactorCompound:      factorCompound,
		IncomeEmployMargin:  incomeEmployMargin,
		IncomeAgreemMargin:  incomeAgreemMargin,
	}
	return guides, nil
}

func NewHealthGuides2018(period Period) (guides *HealthGuides, err error) {
	guides, err = NewHealthGuides(period, HEALTH_PROP_2018_VERSION_MIN,
		HEALTH_PROP_2018_BASIS_MONTHLY_MINIMUM,
		HEALTH_PROP_2018_BASIS_ANNUAL_MAXIMUM,
		HEALTH_PROP_2018_FACTOR_COMPOUND,
		HEALTH_PROP_2018_INCOME_EMPLOY_MARGIN,
		HEALTH_PROP_2018_INCOME_AGREEM_MARGIN)
	return guides, err
}

func NewHealthGuides2017(period Period) (guides *HealthGuides, err error) {
	guides, err = NewHealthGuides(period, HEALTH_PROP_2017_VERSION_MIN,
		HEALTH_PROP_2017_BASIS_MONTHLY_MINIMUM,
		HEALTH_PROP_2017_BASIS_ANNUAL_MAXIMUM,
		HEALTH_PROP_2017_FACTOR_COMPOUND,
		HEALTH_PROP_2017_INCOME_EMPLOY_MARGIN,
		HEALTH_PROP_2017_INCOME_AGREEM_MARGIN)
	return guides, err
}

func NewHealthGuides2016(period Period) (guides *HealthGuides, err error) {
	guides, err = NewHealthGuides(period, HEALTH_PROP_2016_VERSION_MIN,
		HEALTH_PROP_2016_BASIS_MONTHLY_MINIMUM,
		HEALTH_PROP_2016_BASIS_ANNUAL_MAXIMUM,
		HEALTH_PROP_2016_FACTOR_COMPOUND,
		HEALTH_PROP_2016_INCOME_EMPLOY_MARGIN,
		HEALTH_PROP_2016_INCOME_AGREEM_MARGIN)
	return guides, err
}

func NewHealthGuides2015(period Period) (guides *HealthGuides, err error) {
	guides, err = NewHealthGuides(period, HEALTH_PROP_2015_VERSION_MIN,
		HEALTH_PROP_2015_BASIS_MONTHLY_MINIMUM,
		HEALTH_PROP_2015_BASIS_ANNUAL_MAXIMUM,
		HEALTH_PROP_2015_FACTOR_COMPOUND,
		HEALTH_PROP_2015_INCOME_EMPLOY_MARGIN,
		HEALTH_PROP_2015_INCOME_AGREEM_MARGIN)
	return guides, err
}

func NewHealthGuides2014(period Period) (guides *HealthGuides, err error) {
	guides, err = NewHealthGuides(period, HEALTH_PROP_2014_VERSION_MIN,
		HEALTH_PROP_2014_BASIS_MONTHLY_MINIMUM,
		HEALTH_PROP_2014_BASIS_ANNUAL_MAXIMUM,
		HEALTH_PROP_2014_FACTOR_COMPOUND,
		HEALTH_PROP_2014_INCOME_EMPLOY_MARGIN,
		HEALTH_PROP_2014_INCOME_AGREEM_MARGIN)
	return guides, err
}

func NewHealthGuides2013(period Period) (guides *HealthGuides, err error) {
	guides, err = NewHealthGuides(period, HEALTH_PROP_2013_VERSION_MIN,
		HEALTH_PROP_2013_BASIS_MONTHLY_MINIMUM,
		HEALTH_PROP_2013_BASIS_ANNUAL_MAXIMUM,
		HEALTH_PROP_2013_FACTOR_COMPOUND,
		HEALTH_PROP_2013_INCOME_EMPLOY_MARGIN,
		HEALTH_PROP_2013_INCOME_AGREEM_MARGIN)
	if err != nil {
		return nil, err
	}
	if period.Month() <= 7 {
		guides.BasisMonthlyMinimum = HEALTH_TO07_2013_BASIS_MONTHLY_MINIMUM
	}
	return guides, err
}

func NewHealthGuides2012(period Period) (guides *HealthGuides, err error) {
	guides, err = NewHealthGuides(period, HEALTH_PROP_2012_VERSION_MIN,
		HEALTH_PROP_2012_BASIS_MONTHLY_MINIMUM,
		HEALTH_PROP_2012_BASIS_ANNUAL_MAXIMUM,
		HEALTH_PROP_2012_FACTOR_COMPOUND,
		HEALTH_PROP_2012_INCOME_EMPLOY_MARGIN,
		HEALTH_PROP_2012_INCOME_AGREEM_MARGIN)
	return guides, err
}

func NewHealthGuides2011(period Period) (guides *HealthGuides, err error) {
	guides, err = NewHealthGuides(period, HEALTH_PROP_2011_VERSION_MIN,
		HEALTH_PROP_2011_BASIS_MONTHLY_MINIMUM,
		HEALTH_PROP_2011_BASIS_ANNUAL_MAXIMUM,
		HEALTH_PROP_2011_FACTOR_COMPOUND,
		HEALTH_PROP_2011_INCOME_EMPLOY_MARGIN,
		HEALTH_PROP_2011_INCOME_AGREEM_MARGIN)
	return guides, err
}
