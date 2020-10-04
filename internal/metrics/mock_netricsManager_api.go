// Code generated by MockGen. DO NOT EDIT.
// Source: metricsManager.go

// Package metrics is a generated GoMock package.
package metrics

import (
	reflect "reflect"

	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	models "github.com/openshift/assisted-service/models"
	logrus "github.com/sirupsen/logrus"
	time "time"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// ClusterRegistered mocks base method
func (m *MockAPI) ClusterRegistered(clusterVersion string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterRegistered", clusterVersion)
}

// ClusterRegistered indicates an expected call of ClusterRegistered
func (mr *MockAPIMockRecorder) ClusterRegistered(clusterVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRegistered", reflect.TypeOf((*MockAPI)(nil).ClusterRegistered), clusterVersion)
}

// InstallationStarted mocks base method
func (m *MockAPI) InstallationStarted(clusterVersion string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InstallationStarted", clusterVersion)
}

// InstallationStarted indicates an expected call of InstallationStarted
func (mr *MockAPIMockRecorder) InstallationStarted(clusterVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallationStarted", reflect.TypeOf((*MockAPI)(nil).InstallationStarted), clusterVersion)
}

// Duration mocks base method
func (m *MockAPI) Duration(operation string, duration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Duration", operation, duration)
}

// Duration indicates an expected call of Duration
func (mr *MockAPIMockRecorder) Duration(operation, duration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Duration", reflect.TypeOf((*MockAPI)(nil).Duration), operation, duration)
}

// ClusterInstallationFinished mocks base method
func (m *MockAPI) ClusterInstallationFinished(log logrus.FieldLogger, result, clusterVersion string, installationStratedTime strfmt.DateTime) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClusterInstallationFinished", log, result, clusterVersion, installationStratedTime)
}

// ClusterInstallationFinished indicates an expected call of ClusterInstallationFinished
func (mr *MockAPIMockRecorder) ClusterInstallationFinished(log, result, clusterVersion, installationStratedTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterInstallationFinished", reflect.TypeOf((*MockAPI)(nil).ClusterInstallationFinished), log, result, clusterVersion, installationStratedTime)
}

// ReportHostInstallationMetrics mocks base method
func (m *MockAPI) ReportHostInstallationMetrics(log logrus.FieldLogger, clusterVersion string, h *models.Host, previousProgress *models.HostProgressInfo, currentStage models.HostStage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportHostInstallationMetrics", log, clusterVersion, h, previousProgress, currentStage)
}

// ReportHostInstallationMetrics indicates an expected call of ReportHostInstallationMetrics
func (mr *MockAPIMockRecorder) ReportHostInstallationMetrics(log, clusterVersion, h, previousProgress, currentStage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportHostInstallationMetrics", reflect.TypeOf((*MockAPI)(nil).ReportHostInstallationMetrics), log, clusterVersion, h, previousProgress, currentStage)
}
